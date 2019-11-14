package models

import "errors"

type Question struct {
	Description string   `json:"description"`
	Answers     []Answer `json:"answers"`
}

func (q *Question) Validate() (err error) {
	if len(q.Answers) > 8 {
		return errors.New("A question cannot have more than eight answers")
	}
	return nil
}
