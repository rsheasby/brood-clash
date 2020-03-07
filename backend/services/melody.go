package services

import (
	"github.com/olahol/melody"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

var Melody *melody.Melody

func init() {
	Melody = melody.New()

	Melody.HandleMessage(melodyMessageHandler)
}

func melodyMessageHandler(session *melody.Session, bytes []byte) {
	if string(bytes) == "update" {
		_, err := database.GetCurrentQuestionWithAnswers()
		if err != nil {
			session.Write([]byte("Internal Error"))
		}
	} else {
		session.Write([]byte("Invalid Request"))
	}
}
