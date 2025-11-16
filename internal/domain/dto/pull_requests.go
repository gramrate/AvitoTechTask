package dto

import (
	"time"

	"github.com/google/uuid"
)

type PullRequest struct {
	PullRequestID     uuid.UUID   `json:"pull_request_id"`
	PullRequestName   string      `json:"pull_request_name"`
	AuthorID          uuid.UUID   `json:"author_id"`
	Status            string      `json:"status"`
	AssignedReviewers []uuid.UUID `json:"assigned_reviewers"`
	CreatedAt         *time.Time  `json:"-,omitempty"`
	MergedAt          *time.Time  `json:"-,omitempty"`
}

type PullRequestShort struct {
	PullRequestID   uuid.UUID `json:"pull_request_id"`
	PullRequestName string    `json:"pull_request_name"`
	AuthorID        uuid.UUID `json:"author_id"`
	Status          string    `json:"status"`
}

type CreatePRRequest struct {
	PullRequestID   uuid.UUID `json:"pull_request_id"`
	PullRequestName string    `json:"pull_request_name"`
	AuthorID        uuid.UUID `json:"author_id"`
}
type CreatePRResponse struct {
	PR PullRequest `json:"pr"`
}

type MergePRRequest struct {
	PullRequestID uuid.UUID `json:"pull_request_id"`
}
type MergePRResponse struct {
	PR PullRequest `json:"pr"`
}

type ReassignPRRequest struct {
	PullRequestID uuid.UUID `json:"pull_request_id"`
	OldReviewerID uuid.UUID `json:"old_reviewer_id"`
}
type ReassignPRResponse struct {
	PR PullRequest `json:"pr"`
}
