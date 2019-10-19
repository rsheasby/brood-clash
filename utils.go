package main

import (
	"encoding/json"
	"net/http"
)

func EncodeAndReturn(w http.ResponseWriter, status int, object interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(object)
	if status != 200 {
		w.WriteHeader(status)
	}
}
