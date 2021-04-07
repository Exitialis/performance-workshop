package users

import "github.com/pkg/errors"

var (
	userNotFoundError = errors.New("user not found")
	userServiceInternalError = errors.New("user-service 500 error")
)
