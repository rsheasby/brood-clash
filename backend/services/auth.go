package services

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"math/rand"
	"time"
)

var authCode string

func init() {
	rand.Seed(time.Now().UnixNano())
	intCode := rand.Intn(10000)
	// TODO: Remove before releasing
	intCode = 1234
	authCode = fmt.Sprintf("%04d", intCode)
	figure.NewFigure(authCode, "small", true).Print()
}

func IsValidCode(code string) bool {
	return code == authCode
}
