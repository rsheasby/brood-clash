package models

type Question struct {
	BaseModel
	Text         string   `json:"text"`
	Answers      []Answer `json:"answers,omitempty"`
	HasBeenShown bool     `json:"hasBeenShown" readonly:"true"`
}
