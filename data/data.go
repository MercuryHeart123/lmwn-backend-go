package data

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mercuryheart123/lmwn-backend-go/model"
)

func GetData() ([]model.Case, error) {
	response, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err

	}
	var jsonMap model.CasesJSONData
	json.Unmarshal([]byte(responseData), &jsonMap)

	return jsonMap.Data, nil
}
