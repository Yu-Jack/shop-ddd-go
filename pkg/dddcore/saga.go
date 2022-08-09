package dddcore

import (
	"context"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"
)

type SagaData struct {
	Name            string
	Invoke          func(ctx context.Context, value string)
	InvokeKey       string
	Compensation    func(ctx context.Context, value string)
	CompensationKey string
}

type saga struct {
	next     chan int
	steps    []SagaData
	eventBus *EventBus
	ctx      context.Context
}

func NewSaga(ctx context.Context, groupId string, eventBus *EventBus) *saga {
	return &saga{
		next:     make(chan int),
		eventBus: eventBus,
		ctx:      ctx,
	}
}

func (s *saga) AddStep(sd SagaData) error {

	if sd.Invoke == nil {
		return errors.New("invoke should be implemented")
	}

	if sd.InvokeKey == "" {
		return errors.New("InvokeKey should be implemented")
	}

	if (sd.Compensation == nil && sd.CompensationKey == "") ||
		(sd.Compensation != nil && sd.CompensationKey != "") {
		s.steps = append(s.steps, sd)
		return nil
	} else {
		return errors.New("compensatio could be empty or be implemented both")
	}
}

func (s *saga) Execute() error {
	var success int
	i := 0

	for i >= 0 && i < len(s.steps) {
		step := s.steps[i]
		log.Println("start step: ", step.Name)

		readers := []*kafka.Reader{}
		go func() {
			reader := NewReader(step.Name)
			readers = append(readers, reader)
			s.eventBus.SubscribeWithReader(reader, step.InvokeKey, func(value string) {
				step.Invoke(s.ctx, value)
				s.next <- 1
				i++
			})
		}()

		if step.Compensation != nil && step.CompensationKey != "" {
			go func() {
				reader := NewReader(step.Name)
				readers = append(readers, reader)
				s.eventBus.SubscribeWithReader(reader, step.CompensationKey, func(value string) {
					step.Compensation(s.ctx, value)
					s.next <- 2
					i-- // ready to go back execute previous compensation
				})
			}()
		} else if step.Compensation == nil && step.CompensationKey == "" && success == 2 {
			go func() {
				s.next <- 2
				i--
			}()
		}

		success = <-s.next // wait for previous step done

		for _, reader := range readers {
			log.Println("close reader")
			reader.Close()
		}

	}

	return nil
}
