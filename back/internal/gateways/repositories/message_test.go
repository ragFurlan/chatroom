package repository

import (
	entity "chatroom/internal/entities"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage_CreateMessage_Success(t *testing.T) {
	repository := setRepository()

	message := entity.Message{
		UserName: "User 1",
		Message:  "test",
		Room:     "ROOM 1",
	}

	err := repository.CreateMessage(&message)

	assert.NoError(t, err)
}

func TestMessage_CreateMessage_Error(t *testing.T) {
	repository := setRepository()

	message := entity.Message{}

	err := repository.CreateMessage(&message)

	assert.Error(t, err)
}

func TestMessage_GetLatestMessages_Success(t *testing.T) {
	repository := setRepository()

	message := entity.Message{
		UserName: "User 1",
		Message:  "test",
		Room:     "ROOM 1",
	}

	_, err := repository.GetLatestMessages(message.Room, 50)
	assert.NoError(t, err)
}

func setRepository() *MessageRepository {
	dataSourceName := fmt.Sprintf("port=5432 host=localhost user=postgres password=postgres dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	repository := NewMessageRepository(db)
	return repository
}
