package dddcore

import "github.com/google/uuid"

type Event struct {
	UUID      string
	EventName string
	RawData   []byte
}

func NewEvent() Event {
	e := Event{}
	uuid, _ := uuid.NewUUID()
	e.UUID = uuid.String()
	return e
}
