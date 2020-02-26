package models

type Question struct {
	BaseModel
	Text string
	Answers []Answer
}
