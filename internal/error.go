package internal

import "errors"

var (
	ErrDirMissingCode   = errors.New("the provided directory does not contain terraform code")
	ErrFailedToFetch    = errors.New("failed to fetch repository analytics")
	ErrFailedToFindDir  = errors.New("failed to find terraform directory")
	ErrFailedToFindCode = errors.New("failed to find terraform code")
)
