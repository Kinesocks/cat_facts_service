package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	url = "https://catfact.ninja/"
)

// TODO token now in server struct, fix funcs
func BreedsQuery(limit int, token string) []byte {
	urlQuery := url + fmt.Sprintf("breeds?limit=%v", limit)
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlQuery, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accept", "application/json")

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

func GetFactQuery(max_length int, token string) []byte {
	urlQuery := url + fmt.Sprintf("fact?max_length=%v", max_length)
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlQuery, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accept", "application/json")

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

func GetFactsQuery(max_length, limit int, token string) []byte {
	urlQuery := url + fmt.Sprintf("facts?max_length=%v&limit=%v", max_length, limit)
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlQuery, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accept", "application/json")

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
