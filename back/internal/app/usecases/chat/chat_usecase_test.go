package chat_usecase

// import (
// 	producer "chatroom/tests/gateways/producers"
// 	repository "chatroom/tests/gateways/repositories"
// 	usecase "chatroom/tests/usecases"
// 	"errors"
// 	"testing"
//

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"

// )

// var anyError = errors.New("Error")

// func TestChatUseCase_PostMessage_Success(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	controller2 := gomock.NewController(t)
// 	defer controller2.Finish()
// 	controller3 := gomock.NewController(t)
// 	defer controller3.Finish()
// 	controller4 := gomock.NewController(t)
// 	defer controller4.Finish()

// 	userUsecase := usecase.NewMockUser(controller)
// 	botGateway := repository.NewMockBotGateway(controller2)
// 	pubSubProducer := producer.NewMockPubSubGateway(controller3)
// 	messageGateway := repository.NewMockMessageGateway(controller4)
// 	service := NewChatUseCase(botGateway, userUsecase, pubSubProducer, messageGateway)

// 	userID := 1
// 	userName := "User 1"
// 	stockCode := "BA"
// 	value := 100
// 	room := "Room1"
// 	userUsecase.EXPECT().GetUserName(userID).Return(userName, nil)
// 	botGateway.EXPECT().GetStockQuote(stockCode).Return(value, nil)
// 	pubSubProducer.EXPECT().GetSubscribers(room).Return(nil, nil)

// 	err := service.PostMessage(userID, room, stockCode)
// 	assert.NoError(t, err)
// }
