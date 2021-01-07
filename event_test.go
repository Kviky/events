package events

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")
	os.Exit(m.Run())
}

func TestEvent(t *testing.T) {
	eventData := EventData{
		Status: StatusOK,
		Name:   "test",
		Detail: "test detail",
	}
	userData := UserData{
		UserID: "id-test",
		Role:   Admin,
	}
	requestData := RequestData{
		Method: "POST",
		URI:    "/test",
		Body:   `{"name":"test"}`,
	}

	mongoCLient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		Fail(t, err.Error())
	}
	defer mongoCLient.Disconnect(context.TODO())

	repo := NewEventRepostirory(
		mongoCLient,
		os.Getenv("DB_NAME"),
	)

	err = repo.SaveEvent(&userData, &eventData, &requestData)
	NoError(t, err)
}
