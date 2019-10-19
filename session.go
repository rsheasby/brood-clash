package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

var LoginCode string

var SessionToken string

type LoginRequest struct {
	LoginCode string
}

type LoginResponse struct {
	SessionToken string
}

func init() {
	rand.Seed(time.Now().UnixNano())
	GenerateLoginCode()
	GenerateSessionToken()
}

func GenerateLoginCode() {
	code := rand.Intn(10000)
	LoginCode = fmt.Sprintf("%04d", code)
	fmt.Printf("Login code set to: %s\n", LoginCode)
}

func GenerateSessionToken() {
	token, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	SessionToken = token.String()
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
	EncodeAndWriteResponse(w, http.StatusOK, &LoginResponse{SessionToken})
}
