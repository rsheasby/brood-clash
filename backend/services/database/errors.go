package database

import "github.com/rotisserie/eris"

var (
	ErrIDNotFound          = eris.New("Entity with that ID not found in the DB.")
	ErrNoCurrentQuestion   = eris.New("No current question yet. Get unshown question first.")
	ErrAlreadyRevealed     = eris.New("Answer is already revealed.")
	ErrHasAlreadyBeenShown = eris.New("Question has already been shown.")
)
