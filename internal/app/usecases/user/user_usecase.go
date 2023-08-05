package user_usecase

import (
	"chatroom/internal/gateways"
)

type UserUseCase struct {
	UserGateway gateways.UserGateway
}

func NewUserUseCase(userGateway gateways.UserGateway) *UserUseCase {

	return &UserUseCase{
		UserGateway: userGateway,
	}
}

func (uc *UserUseCase) GetUserName(userID int) (string, error) {
	name, err := uc.UserGateway.GetUserName(userID)
	if err != nil {
		return "", err
	}

	return name, nil

}
