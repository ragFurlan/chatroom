package gateway

type UserGateway interface {
	GetUserName(userID string) (string, error)
}
