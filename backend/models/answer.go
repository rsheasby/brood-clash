package models

type Answer struct {
	Text   string `json:"text"`
	Points int    `json:"points"`
}

func (a *Answer) Validate() (err error) {
	return nil
}
