package repository

import (
	entity "chatroom/internal/entities"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	messageRepository := MessageRepository{db: db}
	messageRepository.createTableIfNotExists()
	return &messageRepository
}

func (r *MessageRepository) CreateMessage(message *entity.Message) error {
	insertQuery := fmt.Sprintf("INSERT INTO messages (userName, message, room, timestamp) VALUES ('%s', '%s', '%s', CURRENT_TIMESTAMP)",
		message.UserName, message.Message, message.Room)
	_, err := r.db.Exec(insertQuery)
	return err
}

func (r *MessageRepository) GetLatestMessages(room string, limit int) ([]entity.Message, error) {
	rows, err := r.db.Query("SELECT id, userName, message, room, timestamp FROM messages WHERE room = $1 ORDER BY timestamp DESC LIMIT $2", room, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entity.Message
	for rows.Next() {
		var message entity.Message
		err := rows.Scan(&message.ID, &message.UserName, &message.Room, &message.Message, &message.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (r *MessageRepository) createTableIfNotExists() error {
	sqlFile, err := os.ReadFile("./scripts/create_message.sql")
	if err != nil {
		return fmt.Errorf("Failed to open transaction file: %v", err)
	}

	_, err = r.db.Exec(string(sqlFile))
	return err
}