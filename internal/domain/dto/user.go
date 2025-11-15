package dto

import "github.com/google/uuid"

type User struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	TeamName string    `json:"team_name"`
	IsActive bool      `json:"is_active"`
}

type CreateUserRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	TeamID   uuid.UUID `json:"team_id"`
	IsActive bool      `json:"is_active"`
}
type CreateUserResponse struct {
	User User `json:"user"`
}
type GetUsersPRRequest struct {
	UserID uuid.UUID `json:"user_id" form:"user_id"`
	Status string    `json:"status" form:"status"`
}
type GetUsersPRResponse struct {
	UserID       uuid.UUID           `json:"user_id"`
	PullRequests []*PullRequestShort `json:"pull_requests"`
}

type SetUserActivityRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	IsActive bool      `json:"is_active"`
}
type SetUserActivityResponse struct {
	User User `json:"user"`
}
