package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetFactsQuery(max_length, limit int) []byte {
	urlQuery := url + fmt.Sprintf("facts?max_length=%v&limit=%v", max_length, limit)
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
