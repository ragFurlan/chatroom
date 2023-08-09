package user_usecase

import (
	repository "chatroom/tests/gateways/repositories"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	anyError = errors.New("Error")
)

func TestUserUseCase_GetUserName_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	userGateway := repository.NewMockUserGateway(controller)

	service := NewUserUseCase(userGateway)

	userID := "1"
	userName := "User 1"

	userGateway.EXPECT().GetUserName(userID).Return(userName, nil)

	name, err := service.GetUserName(userID)
	assert.NoError(t, err)
	assert.Equal(t, userName, name)

}

func TestUserUseCase_GetUserName_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	userGateway := repository.NewMockUserGateway(controller)

	service := NewUserUseCase(userGateway)

	userID := "1"
	userName := "User 1"

	userGateway.EXPECT().GetUserName(userID).Return(userName, anyError)

	_, err := service.GetUserName(userID)
	assert.Error(t, err)

}
