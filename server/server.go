package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var port = 8888

type Server struct {
	port     int
	ApiToken string
}

func NewServer() *http.Server {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	NewServer := &Server{
		port:     port,
		ApiToken: os.Getenv("X-CSRF-TOKEN"),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(),
	}

	return server
}
