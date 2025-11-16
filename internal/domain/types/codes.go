package types

type ErrorCode string

const (
	ErrorCodeTeamExists  ErrorCode = "TEAM_EXISTS"
	ErrorCodePRExists    ErrorCode = "PR_EXISTS"
	ErrorCodePRMerged    ErrorCode = "PR_MERGED"
	ErrorCodeNotAssigned ErrorCode = "NOT_ASSIGNED"
	ErrorCodeNoCandidate ErrorCode = "NO_CANDIDATE"
	ErrorCodeNotFound    ErrorCode = "NOT_FOUND"

	ErrorCodeBadRequest    ErrorCode = "BAD_REQUEST"
	ErrorCodeInternalError ErrorCode = "INTERNAL_ERROR"
	ErrorCodeValidation    ErrorCode = "VALIDATION_ERROR"
)
