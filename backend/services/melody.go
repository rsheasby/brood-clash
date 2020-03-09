package services

import (
	"log"

	"github.com/davecgh/go-spew/spew"
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
		q, err := database.GetCurrentQuestionWithAnswers()
		spew.Dump(q)
		if err != nil {
			session.Write([]byte("Internal Error"))
			log.Println(err)
		}
	} else {
		session.Write([]byte("Invalid Request"))
	}
}
