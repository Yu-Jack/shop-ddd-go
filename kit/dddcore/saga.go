package dddcore

import (
	"context"
	"errors"
	"fmt"

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
	reader   *kafka.Reader
	eventBus *EventBus
	ctx      context.Context
}

func NewSaga(ctx context.Context, groupId string, eventBus *EventBus) *saga {
	return &saga{
		next:     make(chan int),
		reader:   NewReader(groupId),
		eventBus: eventBus,
		ctx:      ctx,
	}
}

func (s *saga) AddStep(sd SagaData) {
	s.steps = append(s.steps, sd)
}

func (s *saga) Execute() error {
	defer s.reader.Close()

	for i := 0; i < len(s.steps); i++ {
		step := s.steps[i]
		go func() {
			s.eventBus.SubscribeWithReader(s.reader, step.InvokeKey, func(value string) {
				step.Invoke(s.ctx, value)
			})
			s.next <- 1
		}()
		go func() {
			s.eventBus.SubscribeWithReader(s.reader, step.CompensationKey, func(value string) {
				step.Compensation(s.ctx, value)
			})
			s.next <- 0
		}()
		success := <-s.next // wait for previous step done
		if success != 1 {
			msg := fmt.Sprintf("saga failed in step-%d: %s", i, step.Name)
			return errors.New(msg)
		}
	}

	return nil
}
