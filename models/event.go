package models

import (
	"time"
)

type Event struct {
	ID         uint `gorm:"primarykey"`
	EventID    int
	Created    time.Time
	SystemName string
	IsIncident bool `gorm:"not null;default:false;"`
	Message    string
}

type EventWithoutMessage struct {
	ID         uint
	EventID    int
	Created    time.Time
	SystemName string
	IsIncident bool
}

func (e Event) RemoveMessage() EventWithoutMessage {
	return EventWithoutMessage{
		ID:         e.ID,
		EventID:    e.EventID,
		Created:    e.Created,
		SystemName: e.SystemName,
		IsIncident: e.IsIncident,
	}
}
