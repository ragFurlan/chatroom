package chat_usecase

import (
	entity "chatroom/internal/entities"
	producer "chatroom/tests/gateways/producers"
	repository "chatroom/tests/gateways/repositories"
	usecase "chatroom/tests/usecases"
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	userUsecase    *usecase.MockUser
	botGateway     *repository.MockBotGateway
	pubSubProducer *producer.MockPubSubGateway
	messageGateway *repository.MockMessageGateway
	service        *ChatUseCase
	anyError       = errors.New("Error")
)

func TestChatUseCase_PostMessage_GetUserName_Error(t *testing.T) {
	setControllers(t)

	userID := 1
	userName := "User 1"
	stockCode := "BA"
	room := "Room1"

	userUsecase.EXPECT().GetUserName(userID).Return(userName, anyError)

	err := service.PostMessage(userID, room, stockCode)
	assert.Error(t, err)
}

func TestChatUseCase_PostMessage_GetStockQuote_Error(t *testing.T) {
	setControllers(t)

	userID := 1
	userName := "User 1"
	stockCode := "BA"
	value := 100.5
	room := "Room1"

	userUsecase.EXPECT().GetUserName(userID).Return(userName, nil)
	botGateway.EXPECT().GetStockQuote(stockCode).Return(value, anyError)

	err := service.PostMessage(userID, room, stockCode)
	assert.Error(t, err)
}

func TestChatUseCase_PostMessage_GetStockQuote_Zero_Error(t *testing.T) {
	setControllers(t)

	userID := 1
	userName := "User 1"
	stockCode := "BA"
	value := 0.0
	room := "Room1"

	userUsecase.EXPECT().GetUserName(userID).Return(userName, nil)
	botGateway.EXPECT().GetStockQuote(stockCode).Return(value, nil)

	err := service.PostMessage(userID, room, stockCode)
	assert.Error(t, err)
}

func TestChatUseCase_PostMessage_GetSubscribers_Success(t *testing.T) {
	setControllers(t)

	userID := 1
	userName := "User 1"
	stockCode := "BA"
	value := 100.5
	room := "Room1"
	var ch []chan string
	var chSubscriber chan string
	message := entity.Message{
		UserName: "User 1",
		Message:  "BA quote is $100.5 per share",
		Room:     "Room1",
	}

	jsonBytes, err := json.Marshal(message)

	userUsecase.EXPECT().GetUserName(userID).Return(userName, nil)
	botGateway.EXPECT().GetStockQuote(stockCode).Return(value, nil)
	pubSubProducer.EXPECT().GetSubscribers(room).Return(ch, false)
	pubSubProducer.EXPECT().Subscribe(room).Return(chSubscriber)
	pubSubProducer.EXPECT().Publish(room, string(jsonBytes)).Return()

	err = service.PostMessage(userID, room, stockCode)
	assert.NoError(t, err)
}

func TestChatUseCase_PostMessage_Success(t *testing.T) {
	setControllers(t)

	userID := 1
	userName := "User 1"
	stockCode := "BA"
	value := 100.5
	room := "Room1"
	var ch []chan string

	message := entity.Message{
		UserName: "User 1",
		Message:  "BA quote is $100.5 per share",
		Room:     "Room1",
	}

	jsonBytes, err := json.Marshal(message)

	userUsecase.EXPECT().GetUserName(userID).Return(userName, nil)
	botGateway.EXPECT().GetStockQuote(stockCode).Return(value, nil)
	pubSubProducer.EXPECT().GetSubscribers(room).Return(ch, true)
	pubSubProducer.EXPECT().Publish(room, string(jsonBytes)).Return()

	err = service.PostMessage(userID, room, stockCode)
	assert.NoError(t, err)
}

func TestChatUseCase_GetMessage_Success(t *testing.T) {
	setControllers(t)

	var ch []chan string
	room := "Room1"

	pubSubProducer.EXPECT().GetSubscribers(room).Return(ch, true)
	messageGateway.EXPECT().GetLatestMessages(room, 50).Return([]entity.Message{}, nil)

	_, err := service.GetMessages(room)
	assert.NoError(t, err)
}

func TestChatUseCase_GetMessage_GetSubscribers_Error(t *testing.T) {
	setControllers(t)
	var ch []chan string
	room := "Room1"

	pubSubProducer.EXPECT().GetSubscribers(room).Return(ch, false)

	_, err := service.GetMessages(room)
	assert.Error(t, err)
}

func TestChatUseCase_GetMessage_GetLatestMessages_Error(t *testing.T) {
	setControllers(t)

	var ch []chan string
	room := "Room1"

	pubSubProducer.EXPECT().GetSubscribers(room).Return(ch, true)
	messageGateway.EXPECT().GetLatestMessages(room, 50).Return([]entity.Message{}, anyError)

	_, err := service.GetMessages(room)
	assert.Error(t, err)
}

func setControllers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	controller2 := gomock.NewController(t)
	defer controller2.Finish()
	controller3 := gomock.NewController(t)
	defer controller3.Finish()
	controller4 := gomock.NewController(t)
	defer controller4.Finish()

	userUsecase = usecase.NewMockUser(controller)
	botGateway = repository.NewMockBotGateway(controller2)
	pubSubProducer = producer.NewMockPubSubGateway(controller3)
	messageGateway = repository.NewMockMessageGateway(controller4)
	service = NewChatUseCase(botGateway, userUsecase, pubSubProducer, messageGateway)
}
