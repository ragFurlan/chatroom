package producer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubSub_Subscribe_Success(t *testing.T) {
	service := NewPubSub()
	ch := service.Subscribe("ROOM2")

	assert.NotEqual(t, ch, nil)

}

func TestPubSub_Subscribe_Error(t *testing.T) {
	service := NewPubSub()
	ch := service.Subscribe("")

	assert.NotEqual(t, ch, nil)

}

func TestPubSub_Publish_Success(t *testing.T) {
	service := NewPubSub()
	service.Publish("ROOM2", "test message")
}

func TestPubSub_GetSubscribers_Success(t *testing.T) {
	service := NewPubSub()
	_, bool := service.GetSubscribers("ROOM2")
	assert.Equal(t, bool, false)
}
