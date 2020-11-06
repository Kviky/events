package events

import (
	"context"
	"time"

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

func NewEventRepostirory(dbURI, dbName string) *EventRepository {
	logger := log.New().WithField("repository", "events")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		logger.Fatal(err.Error())
	}

	collection := mongoClient.Database(dbName).Collection("events")

	return &EventRepository{
		Name:        dbName,
		mongoClient: mongoClient,
		collection:  collection,
		logger:      logger,
	}
}

func (r *EventRepository) SaveEvent(userData *UserData, eventData *EventData) error {

	event := &Event{}
	event.SetEventData(eventData)
	event.SetUserData(userData)
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
