package main

import (
	chat "chatroom/internal/app/usecases/chat"
	controller "chatroom/internal/controllers/http/handlers"
	repository "chatroom/internal/repositories"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

)

var (
	handlerChat *controller.HTTPHandler
)

func main() {
	urlStock := os.Getenv("URL_STOCK")

	stockBotRepository := repository.NewStockBotGateway(urlStock)
	chatUsecase := chat.NewChatUseCase(stockBotRepository)
	handlerChat = controller.NewHTTPHandler(*chatUsecase)
	StartServer()

}

func StartServer() {
	router := RegisterRoutes(handlerChat)
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func RegisterRoutes(handlerChat *controller.HTTPHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stock={stock_code}", handlerChat.PostMessageHandler)

	return router
}
