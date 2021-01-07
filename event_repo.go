package events

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventRepository struct {
	Name        string
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewEventRepostirory(mongoClient *mongo.Client, dbName string) *EventRepository {

	collection := mongoClient.Database(dbName).Collection("events")

	return &EventRepository{
		Name:        dbName,
		mongoClient: mongoClient,
		collection:  collection,
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
		return err
	}

	return nil
}
