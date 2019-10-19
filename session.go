package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var LoginCode string

type LoginRequest struct {
	LoginCode string
}

type LoginResponse struct {
	SessionToken string
}

func init() {
	rand.Seed(time.Now().UnixNano())
	GenerateLoginCode()
}

func GenerateLoginCode() {
	code := rand.Intn(10000)
	LoginCode = fmt.Sprintf("%04d", code)
	fmt.Printf("Login code set to: %s\n", LoginCode)
}

func Login(w http.ResponseWriter, r *http.Request) {
	req := new(LoginRequest)
	err := DecodeRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO: Log Juan's mom
		return
	}
	if req.LoginCode != LoginCode {
		w.WriteHeader(http.StatusUnauthorized)
		// TODO: Log Juan's mom
		return
	}
	EncodeAndWriteResponse(w, http.StatusOK, &LoginResponse{"lesessiontoken"})
}
