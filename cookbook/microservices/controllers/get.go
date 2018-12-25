package controllers

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) GetValue(UserDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		value := "default"
		if !UserDefault {
			value = c.storage.Get()
		}
		w.WriteHeader(http.StatusOK)
		p := Payload{Value: value}
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}
