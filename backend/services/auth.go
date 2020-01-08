package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/common-nighthawk/go-figure"
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

func ValidateLoginCode(code string) (isValid bool) {
	return code == loginCode
}
