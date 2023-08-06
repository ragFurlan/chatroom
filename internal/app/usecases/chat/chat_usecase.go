// chat_usecase.go
package chat_usecase

import (
	entity "chatroom/Internal/entities"
	user_usecase "chatroom/internal/app/usecases/user"
	gateway "chatroom/internal/gateways"
	"encoding/json"
	"fmt"
	"time"
)

var shapeTime = "02/01/2006 15:04:05"

type ChatUseCase struct {
	UserUsecase    user_usecase.UserUseCase
	BotGateway     gateway.BotGateway
	PubSubProducer gateway.PubSubGateway
	//MessageRepository MessageRepository
}

func NewChatUseCase(botGateway gateway.BotGateway,
	userUsecase user_usecase.UserUseCase,
	pubSubProducer gateway.PubSubGateway) *ChatUseCase {
	return &ChatUseCase{
		UserUsecase:    userUsecase,
		BotGateway:     botGateway,
		PubSubProducer: pubSubProducer,
		//MessageRepository: messageRepo,
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
		User:      userName,
		Message:   fmt.Sprintf("%s quote is $%v per share", stockCode, value),
		Timestamp: time.Now(),
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("Error transforming message into JSON: %v", err)
	}

	uc.PubSubProducer.Publish(room, string(jsonBytes))

	// message := entity.Message{
	// 	User:      name,
	// 	Message:   fmt.Sprintf("%s.US quote is $%.2f per share", stockCode, value),
	// 	Timestamp: time.Now(),
	// }

	// 	if err := uc.MessageRepository.SaveMessage(&message); err != nil {
	// 	return err
	// }

	return nil

}

func (uc *ChatUseCase) GetMessages(room string) ([]string, error) {
	subscribers, found := uc.PubSubProducer.GetSubscribers(room)
	if !found {
		return nil, fmt.Errorf("There are no posts in this topic")
	}

	err := uc.readMessages(subscribers)
	if err != nil {
		return nil, err
	}

	// Retornar do MySQL todas as mensagens
	return nil, nil

}

func (uc *ChatUseCase) readMessages(subscribers []chan string) error {
	var message entity.Message

	for _, subscription := range subscribers {
		jsonMessage := <-subscription
		err := json.Unmarshal([]byte(jsonMessage), &message)
		if err != nil {
			return fmt.Errorf("Error reading messages: %v", err)
		}

		// TODO: salvar a mensagem no mysql

		// time := message.Timestamp.Format(shapeTime)
		// msg := fmt.Sprintf("%s - %s - %s", message.User, time, message.Message)
		// messages := append(messages, msg)

	}

	return nil
}
