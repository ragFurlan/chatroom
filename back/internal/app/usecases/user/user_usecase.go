package user_usecase

import (
	gateway "chatroom/internal/gateways"
)

type UserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUserUseCase(userGateway gateway.UserGateway) *UserUseCase {
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