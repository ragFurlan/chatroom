package repository

import (
	entity "chatroom/internal/entities"
	"fmt"
	"log"
)

type UserGateway struct{}

func NewUserGateway() *UserGateway {
	return &UserGateway{}
}

func (r UserGateway) GetUserName(userID string) (string, error) {
	log.Printf("service: GetUserName - userID: %v ", userID)
	user, exist := Users[userID]
	if !exist {
		log.Println("service: GetUserName - That user not exist")
		return "", fmt.Errorf("That user not exist")
	}
	return user.Username, nil
}

var Users = map[string]entity.User{
	"1": {ID: "1", Username: "user1", Password: "test123"},
	"2": {ID: "2", Username: "user2", Password: "test123"},
}
