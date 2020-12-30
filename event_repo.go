package events

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventRepository struct {
	Name        string
	mongoClient *mongo.Client
	collection  *mongo.Collection
	logger      *log.Entry
}

func NewEventRepostirory(log log.FieldLogger, mongoClient *mongo.Client, dbName string) *EventRepository {

	logger := log.WithField("repository", "events")
	collection := mongoClient.Database(dbName).Collection("events")

	return &EventRepository{
		Name:        dbName,
		mongoClient: mongoClient,
		collection:  collection,
		logger:      logger,
	}
}

func (r *EventRepository) SaveEvent(userData *UserData, eventData *EventData, requestData *RequestData) error {

	event := &Event{}
	event.SetEventData(eventData)
	event.SetUserData(userData)
	event.SetRequestData(requestData)
	event.SetTimeNow()

	// https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertOne

	_, err := r.collection.InsertOne(context.TODO(), event, options.InsertOne())
	if err != nil {
		r.logger.
			WithError(err).
			WithField("userid", event.UserData.UserID).
			WithField("event-name", event.EventData.Name).
			Error("failed to insert event")
		return err
	}

	r.logger.
		WithField("userid", event.UserData.UserID).
		Info("event has been saved")

	return nil
}
