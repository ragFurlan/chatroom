// chat_usecase.go
package chat_usecase

import (
	user_usecase "chatroom/internal/app/usecases/user"
	"chatroom/internal/gateways"
	"fmt"
)

type ChatUseCase struct {
	UserUsecase user_usecase.UserUseCase
	BotGateway  gateways.BotGateway
	//PubSubProducer    gateways.PubSubProducer
	//MessageRepository MessageRepository
}

func NewChatUseCase(botGateway gateways.BotGateway, userUsecase user_usecase.UserUseCase) *ChatUseCase {
	return &ChatUseCase{
		UserUsecase: userUsecase,
		BotGateway:  botGateway,
		//PubSubProducer:    pubSubProducer,
		//MessageRepository: messageRepo,
	}
}

func (uc *ChatUseCase) PostMessage(userID int, stockCode string) error {
	_, err := uc.UserUsecase.GetUserName(userID)
	if err != nil {
		return err
	}

	// message := entities.Message{
	// 	User:      user.Username,
	// 	Content:   content,
	// 	Timestamp: time.Now(),
	// }

	// if err := uc.MessageRepository.SaveMessage(&message); err != nil {
	// 	return err
	// }

	value, err := uc.BotGateway.GetStockQuote(stockCode)
	if err != nil {
		return err
	}

	if value == 0 {
		return fmt.Errorf("This stock code does not exist")
	}

	// Send the stock quote back to the chatroom using Cloud Pub/Sub
	// botMessage := fmt.Sprintf("%s quote is $%s per share", strings.TrimPrefix(content, "/stock="), quote)
	// if err := uc.PubSubProducer.SendMessageToTopic(botMessage); err != nil {
	// 	return err
	// }

	return nil

}

// func (uc *ChatUseCase) GetLatestMessages() ([]entities.Message, error) {
// 	messages, err := uc.MessageRepository.GetLatestMessages(50)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return messages, nil
// }
