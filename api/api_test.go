package api

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetBreeds(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	token := os.Getenv("X-CSRF-TOKEN")

	input := 6
	var res map[string]interface{}
	json.Unmarshal(BreedsQuery(input, token), &res)
	if _, ok := res["data"]; !ok {
		t.Error("Didn't get \"data\" field from responce")
	}

	data := res["data"].([]interface{})
	if len(data) == input-1 {
		t.Error("Didn't get breeds from api")
	}
}
