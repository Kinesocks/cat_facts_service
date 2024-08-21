package api

import (
	"encoding/json"
	"testing"
)

func TestGetBreeds(t *testing.T) {
	input := 6
	var res map[string]interface{}
	json.Unmarshal(BreedsQuery(input), &res)
	if _, ok := res["data"]; !ok {
		t.Error("Didn't get \"data\" field from responce")
	}
	data := res["data"].([]interface{})
	if len(data) == input-1 {
		t.Error("Didn't get breeds from api")
	}
}
