package services

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
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
	fmt.Println("Your login code is:")
	myFigure := figure.NewFigure(loginCode, "", true)
	myFigure.Print()
	fmt.Println()
}

func LoginCodeValidator(code string, _ echo.Context) (isValid bool, err error) {
	return code == loginCode, nil
}
