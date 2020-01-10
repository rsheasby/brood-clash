package services

import (
	"backend/models"
)

var updateListeners []chan []models.Question

func init() {
	updateListeners = make([]chan []models.Question, 0)
}

func RegisterUpdateListener(listener chan []models.Question) {
	updateListeners = append(updateListeners, listener)
}

func UnregisterUpdateListener(listener chan []models.Question) (removed bool) {
	for i := range updateListeners {
		if updateListeners[i] == listener {
			updateListeners = append(updateListeners[:i], updateListeners[i+1:]...)
			return true
		}
	}
	return false
}

func NotifyUpdateListeners() {
	questions := GetQuestions()
	for _, v := range updateListeners {
		v <- questions
	}
}
