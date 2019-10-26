package main

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, object interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(object)
	return
}

func EncodeAndWriteResponse(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(object)
	if status != 200 {
		w.WriteHeader(status)
	}
}
