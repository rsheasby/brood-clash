package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
)

var loginCode string

func init() {
	rand.Seed(time.Now().UnixNano())
	GenerateLoginCode()
}

func GenerateLoginCode() {
	code := rand.Intn(10000)
	loginCode = fmt.Sprintf("%04d", code)
	fmt.Printf("Login code set to: %s\n", loginCode)
}

func LoginCodeValidator(code string, c echo.Context) (isValid bool, err error) {
	return code == loginCode, nil
}
