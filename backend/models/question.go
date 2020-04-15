package models

type Question struct {
	BaseModel
	Text         string
	Answers      []Answer `json:",omitempty"`
	HasBeenShown bool     `json:"-"`
}
