package handler

import (
	chat_usecase "chatroom/internal/app/usecases/chat"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	ChatUseCase chat_usecase.ChatUseCase
}

func NewHTTPHandler(chatUseCase chat_usecase.ChatUseCase) *HTTPHandler {
	return &HTTPHandler{
		ChatUseCase: chatUseCase,
	}
}

func (h *HTTPHandler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("UserID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	room := r.Header.Get("Room")

	vars := mux.Vars(r)
	stockCode, ok := vars["stock_code"]
	if !ok {
		http.Error(w, "Id is missing in parameters", http.StatusMethodNotAllowed)
	}

	err = h.ChatUseCase.PostMessage(userID, room, stockCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// room
