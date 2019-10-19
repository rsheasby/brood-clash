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
	encoder := json.NewEncoder(w)
	encoder.Encode(object)
	if status != 200 {
		w.WriteHeader(status)
	}
}
