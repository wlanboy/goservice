package model

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

/*Event struct*/
type Event struct {
	//gorm.Model
	ID        uint64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT;INDEX"`
	UUID      uuid.UUID  `gorm:"type:uuid;INDEX;NOT NULL" json:"uuid"`
	Name      string     `gorm:"NOT NULL" json:"name"`
	Type      string     `gorm:"NOT NULL" json:"type"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}

/*Validate Event struct*/
func (event *Event) Validate() (string, bool) {

	if event.Name == "" {
		return "name missing", false
	}

	if event.Type == "" {
		return "Event Type missing", false
	}

	return "", true
}

/*SaveEvent struct*/
func SaveEvent(event Event) (string, *Event) {

	if error, ok := event.Validate(); !ok {
		return error, nil
	} else {
		uuid, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("uuid.NewV4 went wrong: %s", err)
		} else {
			event.UUID = uuid
		}
		inserterr := GetDB().Create(&event).Error
		if inserterr != nil {
			return fmt.Sprintf("Event cannot be saved %s", inserterr), nil
		} else {
			return "", &event
		}
	}
}

/*GetEventByID uint*/
func GetEventByID(id uuid.UUID) (string, *Event) {

	event := Event{}
	err := GetDB().Where("uuid = ?", id).First(&event).Error
	if err != nil {
		return fmt.Sprintf("No event found with id %d", id), nil
	}
	return "", &event
}

/*GetAllEvents array*/
func GetAllEvents() (string, []*Event) {

	events := make([]*Event, 0)
	err := GetDB().Find(&events).Error
	if err != nil {
		return "", nil
	}

	return "", events
}
