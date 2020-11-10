package events

import (
	"time"
)

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
		UserID    string `bson:"userid"`
		Role      Role   `bson:"role"`
		CharterID string `bson:"charterid"`
	}

	// EventData is a struct wich has information about event
	EventData struct {
		Status EventStatus `bson:"status"`
		// Name where is the event happened
		Name string `bson:"name"`
		// Detail information of the event
		Detail string `bson:"detail"`
	}

	// Request data has information about request
	RequestData struct {
		// Method is a http method
		Method string `bson:"method"`
		// URI is request URI
		URI string `bson:"uri"`
		// Body is a body from request
		Body string `bson:"body"`
	}

	Event struct {
		UserData    *UserData    `bson:"user_data"`
		EventData   *EventData   `bson:"event_data"`
		RequestData *RequestData `bson:"request_data"`
		CreatedAt   time.Time    `bson:"created_at"`
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

func (e *Event) SetRequestData(request *RequestData) *Event {
	e.RequestData = request
	return e
}

func (e *Event) SetTimeNow() *Event {
	e.CreatedAt = time.Now()
	return e
}
