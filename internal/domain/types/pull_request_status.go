package types

import (
	"fmt"
)

type PullRequestStatus int

const (
	PullRequestStatusOpen PullRequestStatus = iota + 1
	PullRequestStatusMerged
)

// String возвращает строковое представление
func (s PullRequestStatus) String() string {
	switch s {
	case PullRequestStatusOpen:
		return "OPEN"
	case PullRequestStatusMerged:
		return "MERGED"
	default:
		return "UNKNOWN"
	}
}

// FromString создает PullRequestStatus из строки
func FromString(str string) (PullRequestStatus, error) {
	switch str {
	case "OPEN":
		return PullRequestStatusOpen, nil
	case "MERGED":
		return PullRequestStatusMerged, nil
	default:
		return PullRequestStatusOpen, fmt.Errorf("invalid PullRequestStatus: %s", str)
	}
}
