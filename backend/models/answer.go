package models

type Answer struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Points   int    `json:"points"`
	Revealed bool   `json:"revealed"`
}

func (a *Answer) Validate() (err error) {
	return nil
}
