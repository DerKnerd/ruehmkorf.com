package api

import (
	"encoding/json"
	"net/http"
	"ruehmkorf/database"
)

func getSpotMappings(w http.ResponseWriter, r *http.Request) {
	mappings, err := database.Select[database.SpotMapping]("select * from spot_mapping")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(mappings)
}

func replaceSpotMappings(w http.ResponseWriter, r *http.Request) {
	var body []struct {
		Character string `json:"character"`
		English   string `json:"english"`
		German    string `json:"german"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.GetDbMap().Exec("delete from spot_mapping")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mappings := make([]interface{}, len(body))
	for i, item := range body {
		mappings[i] = &database.SpotMapping{
			Character: item.Character,
			English:   item.English,
			German:    item.German,
		}
	}

	err = database.GetDbMap().Insert(mappings...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
