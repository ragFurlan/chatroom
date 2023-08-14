package main

import (
	chat_usecase "chatroom/internal/app/usecases/chat"
	userUsecase "chatroom/internal/app/usecases/user"
	controller "chatroom/internal/controllers/http/handlers"
	producer "chatroom/internal/gateways/producers"
	repository "chatroom/internal/gateways/repositories"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	handlerChat *controller.HTTPHandler
	sub         <-chan string
)

func main() {
	log.Println("Starting server...")

	urlStock := os.Getenv("URL_STOCK")

	// User
	userGateway := repository.NewUserGateway()
	userUseCase := userUsecase.NewUserUseCase(userGateway)

	// Bot
	stockBotRepository := repository.NewStockBotGateway(urlStock)

	// pub/sub
	pubSubProducer := producer.NewPubSub()

	// Repository
	messageRepository := setRepository()

	// Chat
	chatUsecase := chat_usecase.NewChatUseCase(stockBotRepository, userUseCase, pubSubProducer, messageRepository)

	handlerChat = controller.NewHTTPHandler(*chatUsecase)
	startServer()

}

func setRepository() *repository.MessageRepository {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dataSourceName := fmt.Sprintf("port=5432 host=db user=%s password=%s dbname=postgres sslmode=disable ", user, password)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start service: setRepository - method: setRepository")
	repository := repository.NewMessageRepository(db)

	log.Println("finish service: setRepository - method: setRepository")
	return repository
}

func startServer() {
	router := registerRoutes(handlerChat)
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func registerRoutes(handlerChat *controller.HTTPHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/message/stock={stock_code}", handlerChat.PostMessageHandler).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/message", handlerChat.GetMessages).Methods("GET", "POST", "OPTIONS")

	return router
}
