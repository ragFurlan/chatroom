package gateway

import (
	entity "chatroom/internal/entities"
)

type MessageGateway interface {
	CreateMessage(message *entity.Message) error
	GetLatestMessages(room string, limit int) ([]entity.Message, error)
}
