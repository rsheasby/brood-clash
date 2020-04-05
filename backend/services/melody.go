package services

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

var mel *melody.Melody

type UpdateType string

const (
	incorrectAnswer = "incorrectAnswer"
	revealAnswer = "revealAnswer"
	stateUpdate = "stateUpdate"
)

type UpdateMessage struct {
	Type UpdateType
	Update interface{} `json:",omitempty"`
}

func init() {
	mel = melody.New()

	mel.HandleConnect(melodyConnectHandler)
	mel.HandleMessage(melodyMessageHandler)
}

func UpgradeToWebsocket(w http.ResponseWriter, r *http.Request) error {
	return mel.HandleRequest(w, r)
}

func BroadcastIncorrectAnswer() {
	um := UpdateMessage{
		Type:   incorrectAnswer,
		Update: nil,
	}
	msg, err := json.Marshal(um)
	if err != nil {
		return
	}
	mel.Broadcast(msg)
}

func BroadcastRevealAnswer(uuid uuid.UUID) {
	answer, err := database.GetAnswer(uuid)
	if err != nil {
		return
	}
	um := UpdateMessage{
		Type:   revealAnswer,
		Update: answer,
	}
	msg, err := json.Marshal(um)
	if err != nil {
		return
	}
	mel.Broadcast(msg)
}

func BroadcastStateUpdate() {
	mel.Broadcast(getUpdateMsg())
}

func getUpdateMsg() []byte {
	q, err := database.GetCurrentQuestionWithAnswers()
	if err != nil {
		return []byte("{}")
	}
	for i := range q.Answers {
		if !q.Answers[i].Revealed {
			q.Answers[i].Text = ""
			q.Answers[i].Points = 0
		}
	}
	um := UpdateMessage{
		Type:   stateUpdate,
		Update: q,
	}
	msg, err := json.Marshal(um)
	if err != nil {
		return []byte("{}")
	}
	return msg
}

func melodyConnectHandler(session *melody.Session) {
	session.Write(getUpdateMsg())
}

func melodyMessageHandler(session *melody.Session, bytes []byte) {
	if string(bytes) == "update" {
		session.Write(getUpdateMsg())
	} else {
		session.Write([]byte("Invalid Request"))
	}
}
