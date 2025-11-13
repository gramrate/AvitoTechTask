package dto

import (
	"AvitoTechTask/internal/domain/types"
	"time"

	"github.com/google/uuid"
)

type PullRequest struct {
	PullRequestID     uuid.UUID               `json:"pull_request_id"`
	PullRequestName   string                  `json:"pull_request_name"`
	AuthorID          uuid.UUID               `json:"author_id"`
	Status            types.PullRequestStatus `json:"status"`
	AssignedReviewers []uuid.UUID             `json:"assigned_reviewers"`
	CreatedAt         *time.Time              `json:"created_at,omitempty"`
	MergedAt          *time.Time              `json:"merged_at,omitempty"`
}

type PullRequestShort struct {
	PullRequestID   uuid.UUID               `json:"pull_request_id"`
	PullRequestName string                  `json:"pull_request_name"`
	AuthorID        uuid.UUID               `json:"author_id"`
	Status          types.PullRequestStatus `json:"status"`
}
