package repository

import (
	entity "chatroom/Internal/entities"
	"fmt"
)

type UserGateway struct{}

func NewUserGateway() *UserGateway {
	return &UserGateway{}
}

func (r UserGateway) GetUserName(userID int) (string, error) {
	user, exist := Users[userID]
	if !exist {
		return "", fmt.Errorf("That user not exist")
	}
	return user.Username, nil
}

var Users = map[int]entity.User{
	1: {ID: 1, Username: "user1", Password: "test123"},
	2: {ID: 2, Username: "user2", Password: "test123"},
}
