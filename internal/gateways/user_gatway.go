package gateway

type UserGateway interface {
	GetUserName(userID int) (string, error)
}
