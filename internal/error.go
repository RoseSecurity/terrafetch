package internal

import "errors"

var (
	ErrDirMissingCode = errors.New("the provided directory does not contain terraform code")
	ErrFailedToFetch  = errors.New("failed to fetch repository analytics")
)
