package gateways

type UserGateway interface {
	GetUserName(userID int) (string, error)
}
