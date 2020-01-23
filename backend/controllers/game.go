package controllers

import (
	"backend/models"
	"backend/restapi/operations"
	"backend/services"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
)

func init() {
	ControllerConfigList.Register(func(api *operations.BroodClashAPI) {
		api.CreateQuestionsHandler = operations.CreateQuestionsHandlerFunc(CreateQuestionsHandler)
		api.GetQuestionHandler = operations.GetQuestionHandlerFunc(GetQuestionHandler)
		api.GetQuestionsHandler = operations.GetQuestionsHandlerFunc(GetQuestionsHandler)
		api.RevealAnswerHandler = operations.RevealAnswerHandlerFunc(RevealAnswerHandler)
	})
}

func CreateQuestionsHandler(params operations.CreateQuestionsParams, _ interface{}) middleware.Responder {
	for _, v := range params.Questions {
		_, err := services.CreateQuestion(*v)
		if err != nil {
			// TODO: Log this error somehow properly.
			fmt.Println(err)
		}
	}
	return operations.NewCreateQuestionCreated()
}

func GetQuestionHandler(params operations.GetQuestionParams, _ interface{}) middleware.Responder {
	question := services.GetQuestion(params.ID)
	if question == nil {
		return operations.NewGetQuestionNotFound()
	} else {
		return operations.NewGetQuestionOK().WithPayload(question)
	}
}

func GetQuestionsHandler(params operations.GetQuestionsParams, _ interface{}) middleware.Responder {
	output := make([]*models.Question, 0)
	questions := services.GetQuestions()
	for i := range questions {
		output = append(output, &(questions[i]))
	}
	return operations.NewGetQuestionsOK().WithPayload(output)
}

func RevealAnswerHandler(params operations.RevealAnswerParams, _ interface{}) middleware.Responder {
	question := services.GetQuestion(params.QuestionID)
	if question == nil {
		return operations.NewRevealAnswerNotFound()
	}

	for i := range question.Answers {
		if question.Answers[i].ID == params.AnswerID {
			answer := question.Answers[i]
			if answer.Revealed == nil {
				answer.Revealed = new(bool)
			}
			*answer.Revealed = true
			services.NotifyUpdateListeners()
			return operations.NewRevealAnswerNoContent()
		}
	}

	return operations.NewRevealAnswerNotFound()
}
