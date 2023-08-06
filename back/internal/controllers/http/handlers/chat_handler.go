package handler

import (
	chat_usecase "chatroom/internal/app/usecases/chat"
	"encoding/json"
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

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Message posted with success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("UserID")
	_, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	room := r.Header.Get("Room")

	messages, err := h.ChatUseCase.GetMessages(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)

}
