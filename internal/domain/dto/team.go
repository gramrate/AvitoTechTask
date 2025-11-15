package dto

import (
	"github.com/google/uuid"
)

type Team struct {
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}
type TeamMember struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	IsActive bool      `json:"is_active"`
}

type CreateTeamRequest struct {
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}
type CreateTeamResponse struct {
	Team Team `json:"team"`
}

type GetTeamRequest struct {
	TeamName string `json:"team_name"`
}
type GetTeamResponse Team
