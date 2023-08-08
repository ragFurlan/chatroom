package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_CreateMessage_Success(t *testing.T) {
	repository := NewUserGateway()

	name, err := repository.GetUserName(1)

	assert.NoError(t, err)
	assert.Equal(t, "user1", name)
}

func TestUser_CreateMessage_Error(t *testing.T) {
	repository := NewUserGateway()

	_, err := repository.GetUserName(5)

	assert.Error(t, err)
}
