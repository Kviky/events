package events

import "time"

const (
	StatusOK     EventStatus = "OK"
	StatusFailed EventStatus = "FAILED"
)

const (
	Admin   Role = "ADMIN"
	Charter Role = "CHARTER"
)

type (
	// EventStatus is a status of event. OK/FAILED
	EventStatus string

	// Role is a UserRole type. ADMIN/CHARTER
	Role string

	// UserData is a struct wich has information about user
	UserData struct {
		UserID string `bson:"userid"`
		Role   Role   `bson:"role"`
	}

	// EventData is a struct wich has information about event
	EventData struct {
		Status EventStatus `bson:"status"`
		// Name where is the event happened
		Name string `bson:"name"`
		// Detail information of the event
		Detail string `bson:"detail"`
	}

	Event struct {
		UserData  *UserData  `bson:"user_data"`
		EventData *EventData `bson:"event_data"`
		CreatedAt time.Time  `bson:"created_at"`
	}
)

func (e *Event) SetUserData(user *UserData) *Event {
	e.UserData = user
	return e
}

func (e *Event) SetEventData(event *EventData) *Event {
	e.EventData = event
	return e
}

func (e *Event) SetTimeNow() *Event {
	e.CreatedAt = time.Now()
	return e
}
