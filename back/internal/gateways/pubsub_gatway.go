package gateway

type PubSubGateway interface {
	Subscribe(topic string) chan string
	Publish(topic, message string)
	GetSubscribers(room string) ([]chan string, bool)
}
