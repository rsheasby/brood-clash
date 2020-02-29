package services

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type WebSocketManager struct {
	sync.Mutex
	channels []chan uuid.UUID
}

var wsm WebSocketManager

func init() {
	wsm.channels = make([]chan uuid.UUID, 0)
}

func RegisterForAnswerNotification() chan uuid.UUID {
	wsm.Lock()
	defer wsm.Unlock()
	c := make(chan uuid.UUID)
	wsm.channels = append(wsm.channels, c)
	return c
}

func UnregisterAnswerNotification(c chan uuid.UUID) {
	wsm.Lock()
	defer wsm.Unlock()
	for i := range wsm.channels {
		if wsm.channels[i] == c {
			wsm.channels[i] = wsm.channels[len(wsm.channels)-1]
			wsm.channels = wsm.channels[:len(wsm.channels)-1]
			return
		}
	}
}

func SendAnswerNotification(answerID uuid.UUID) {
	wsm.Lock()
	defer wsm.Unlock()

	for i := range wsm.channels {
		go func(c chan uuid.UUID) {
			timeout := time.NewTimer(2 * time.Second)
			select {
			case <-timeout.C:
				go UnregisterAnswerNotification(c)
			case c <- answerID:
			}
			timeout.Stop()
		}(wsm.channels[i])
	}
}
