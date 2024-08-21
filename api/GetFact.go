package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetFactQuery(max_length int) []byte {
	urlQuery := url + fmt.Sprintf("fact?max_length=%v", max_length)
	client := &http.Client{}
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	req, err := http.NewRequest("GET", urlQuery, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accept", "application/json")

	token := os.Getenv("X-CSRF-TOKEN")
	req.Header.Set("X-CSRF-TOKEN", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return respBody
}
