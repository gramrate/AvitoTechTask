package errorz

import "errors"

var (
	ErrTeamNotFound        = errors.New("team not found")
	ErrTeamNameAlreadyUsed = errors.New("team name already used")
)
