package chat_usecase

import (
	user_usecase "chatroom/internal/app/usecases/user"
	entity "chatroom/internal/entities"
	gateway "chatroom/internal/gateways"
	"encoding/json"
	"fmt"
)

var (
	shapeTime       = "02/01/2006 15:04:05"
	limitMessageGet = 50
)

type ChatUseCase struct {
	UserUsecase       user_usecase.User
	BotGateway        gateway.BotGateway
	PubSubProducer    gateway.PubSubGateway
	MessageRepository gateway.MessageGateway
}

type Chat interface {
	PostMessage(userID int, room, stockCode string) error
	GetMessages(room string) ([]entity.Message, error)
}

func NewChatUseCase(botGateway gateway.BotGateway,
	userUsecase user_usecase.User,
	pubSubProducer gateway.PubSubGateway,
	messageRepository gateway.MessageGateway) *ChatUseCase {
	return &ChatUseCase{
		UserUsecase:       userUsecase,
		BotGateway:        botGateway,
		PubSubProducer:    pubSubProducer,
		MessageRepository: messageRepository,
	}
}

func (uc *ChatUseCase) PostMessage(userID int, room, stockCode string) error {
	userName, err := uc.UserUsecase.GetUserName(userID)
	if err != nil {
		return err
	}

	value, err := uc.BotGateway.GetStockQuote(stockCode)
	if err != nil {
		return err
	}

	if value == 0 {
		return fmt.Errorf("This stock code does not exist")
	}

	_, found := uc.PubSubProducer.GetSubscribers(room)
	if !found {
		uc.PubSubProducer.Subscribe(room)
	}

	message := entity.Message{
		UserName: userName,
		Message:  fmt.Sprintf("%s quote is $%v per share", stockCode, value),
		Room:     room,
	}

	jsonBytes, _ := json.Marshal(message)
	uc.PubSubProducer.Publish(room, string(jsonBytes))

	return nil
}

func (uc *ChatUseCase) GetMessages(room string) ([]entity.Message, error) {
	subscribers, found := uc.PubSubProducer.GetSubscribers(room)
	if !found {
		return nil, fmt.Errorf("There are no posts in this topic")
	}

	err := uc.readMessages(subscribers)
	if err != nil {
		return nil, err
	}

	messages, err := uc.MessageRepository.GetLatestMessages(room, limitMessageGet)
	if err != nil {
		return nil, fmt.Errorf("Error add message to database: %v", err)
	}

	return messages, nil

}

func (uc *ChatUseCase) readMessages(subscribers []chan string) error {
	var message entity.Message
	for _, ch := range subscribers {
		select {
		case jsonMessage := <-ch:
			err := json.Unmarshal([]byte(jsonMessage), &message)
			if err != nil {
				return fmt.Errorf("Error reading messages: %v", err)
			}

			err = uc.MessageRepository.CreateMessage(&message)
			if err != nil {
				return fmt.Errorf("Error add message to database: %v", err)
			}
		default:
			continue
		}
	}

	return nil
}
