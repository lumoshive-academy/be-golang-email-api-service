package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lumoshive-academy/be-golang-email-api-service/handler"
	"github.com/lumoshive-academy/be-golang-email-api-service/middleware"
	"github.com/lumoshive-academy/be-golang-email-api-service/utils"
)

func main() {
	// read config
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Println("error :", err)
	}
	handler := handler.NewHandlerEmail(config)

	router := chi.NewRouter()
	router.With(middleware.ApiKeyMiddleware).Post("/send-email", handler.SendEmail)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
