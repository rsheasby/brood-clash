package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/rsheasby/brood-clash/backend/config"
)

var authCode string

func init() {
	rand.Seed(time.Now().UnixNano())
	var intCode int
	if config.RandomizeLoginCode {
		intCode = rand.Intn(10000)
	} else {
		intCode = 1234
	}
	authCode = fmt.Sprintf("%04d", intCode)
	fmt.Println("Login Code:")
	figure.NewFigure(authCode, "big", true).Print()
	fmt.Println()
}

func IsValidCode(code string) bool {
	return code == authCode
}
