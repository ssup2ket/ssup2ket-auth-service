package service

import (
	"fmt"

	"github.com/ssup2ket/ssup2ket-auth-service/internal/domain/repo"
)

// Error
var (
	// Common
	ErrServerErr error = fmt.Errorf("server error")

	// Auth
	ErrUnauthorized error = fmt.Errorf("unauthorized")

	// Repository
	ErrRepoNotFound    error = fmt.Errorf("repo resource not found")
	ErrRepoConflict    error = fmt.Errorf("repo conflict")
	ErrRepoServerError error = fmt.Errorf("repo server error")
)

func getReturnErr(err error) error {
	switch err {
	case repo.ErrNotFound:
		return ErrRepoNotFound
	case repo.ErrConflict:
		return ErrRepoConflict
	case repo.ErrServerError:
		return ErrRepoServerError
	}
	return ErrServerErr
}
