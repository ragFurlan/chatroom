package handler

import (
	chat_usecase "chatroom/internal/app/usecases/chat"
	"encoding/json"
	"log"
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
		UserID string    `json:"UserId"`
		Room   string `json:"room"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("service: PostMessageHandler - userID: %v ", requestBody)

	vars := mux.Vars(r)
	stockCode, ok := vars["stock_code"]
	if !ok {
		log.Println("service: PostMessageHandler - Id is missing in parameters")
		http.Error(w, "Id is missing in parameters", http.StatusMethodNotAllowed)
	}

	err = h.ChatUseCase.PostMessage( requestBody.UserID, requestBody.Room, stockCode)
	if err != nil {
		log.Printf("service: PostMessageHandler - method: PostMessage - error: %v ", err)
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

	log.Printf("service: GetMessages - requestBody: %v ", requestBody)

	messages, err := h.ChatUseCase.GetMessages(requestBody.Room)
	if err != nil {
		log.Printf("service: GetMessages - method: GetMessages  error: %v ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(messages)

}
