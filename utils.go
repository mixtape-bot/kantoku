package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}

func (k *Kantoku) createJsonResponse(w http.ResponseWriter, data any, success bool) {
	k.createJson(w, Response{
		Data:    data,
		Success: success,
	})
}

func (k *Kantoku) createJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		k.Logger.Errorf("Failed to create JSON response: %s", err)
	}
}
