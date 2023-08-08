package handler

import (
	chat_usecase "chatroom/internal/app/usecases/chat"
	"encoding/json"
	"net/http"

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
	var requestBody struct {
		UserID int    `json:"UserId"`
		Room   string `json:"room"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	stockCode, ok := vars["stock_code"]
	if !ok {
		http.Error(w, "Id is missing in parameters", http.StatusMethodNotAllowed)
	}

	//userID, _ := strconv.Atoi(requestBody.UserID)
	err = h.ChatUseCase.PostMessage(requestBody.UserID, requestBody.Room, stockCode)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Room string `json:"room"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	messages, err := h.ChatUseCase.GetMessages(requestBody.Room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(messages)

}
