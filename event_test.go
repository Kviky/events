package events

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/stretchr/testify/assert"
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

	repo := NewEventRepostirory(
		os.Getenv("DB_URI"),
		os.Getenv("DB_NAME"),
	)

	err := repo.SaveEvent(&userData, &eventData)
	NoError(t, err)
}
