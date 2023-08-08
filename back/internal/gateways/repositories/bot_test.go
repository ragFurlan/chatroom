package repository

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBot_Subscribe_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := NewStockBotGateway("./")
	_, err := service.GetStockQuote("BA")

	assert.NoError(t, err)
}
