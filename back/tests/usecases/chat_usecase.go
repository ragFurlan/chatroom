// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/usecases/chat/chat_usecase.go

// Package usecase is a generated GoMock package.
package usecase

import (
	entity "chatroom/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChat is a mock of Chat interface.
type MockChat struct {
	ctrl     *gomock.Controller
	recorder *MockChatMockRecorder
}

// MockChatMockRecorder is the mock recorder for MockChat.
type MockChatMockRecorder struct {
	mock *MockChat
}

// NewMockChat creates a new mock instance.
func NewMockChat(ctrl *gomock.Controller) *MockChat {
	mock := &MockChat{ctrl: ctrl}
	mock.recorder = &MockChatMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChat) EXPECT() *MockChatMockRecorder {
	return m.recorder
}

// GetMessages mocks base method.
func (m *MockChat) GetMessages(room string) ([]entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", room)
	ret0, _ := ret[0].([]entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockChatMockRecorder) GetMessages(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockChat)(nil).GetMessages), room)
}

// PostMessage mocks base method.
func (m *MockChat) PostMessage(userID int, room, stockCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMessage", userID, room, stockCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMessage indicates an expected call of PostMessage.
func (mr *MockChatMockRecorder) PostMessage(userID, room, stockCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockChat)(nil).PostMessage), userID, room, stockCode)
}
