package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Kinesocks/cat_facts_service/api"
)

func GetBreeds(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	var query map[string]interface{}
	err = json.Unmarshal(b, &query)
	if err != nil {
		log.Fatal(err)
	}

	_, ok := query["limit"]

	if !ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	lim := int(query["limit"].(float64))

	resp := api.BreedsQuery(lim)

	var apiresp map[string]interface{}
	err = json.Unmarshal(resp, &apiresp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiresp["data"])
}
