package producer

import (
	"log"
	"sync"

)

type PubSub struct {
	subscribers map[string][]chan string
	mu          sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) chan string {
	log.Printf("service: Subscribe - message: %v", topic)
	if topic == "" {
		return nil
	}
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	ch := make(chan string, 10)

	delete(ps.subscribers, topic)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)

	return ch
}

func (ps *PubSub) Publish(topic, message string) {
	log.Printf("service: Publish - room: %v - message: %v", topic, message)
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	subscribers, found := ps.subscribers[topic]
	if !found {
		return
	}

	for _, ch := range subscribers {
		go func(c chan string) {
			c <- message
		}(ch)
	}
}

func (ps *PubSub) GetSubscribers(room string) ([]chan string, bool) {
	log.Printf("service: GetSubscribers - room: %v ", room)
	subscribers, found := ps.subscribers[room]
	if !found {
		return nil, false
	}
	return subscribers, true

}
