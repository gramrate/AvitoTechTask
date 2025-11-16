package errorz

import "errors"

var (
	ErrPRNotFound        = errors.New("pull request not found")
	ErrPRMerged          = errors.New("PR_MERGED: cannot modify merged pull request")
	ErrNotAssigned       = errors.New("NOT_ASSIGNED: reviewer is not assigned to this pull request")
	ErrNoCandidate       = errors.New("NO_CANDIDATE: no available reviewers found")
	ErrPRNameAlreadyUsed = errors.New("pull request name already used")
)
