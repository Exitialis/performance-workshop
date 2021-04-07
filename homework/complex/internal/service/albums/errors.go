package albums

import "github.com/pkg/errors"

var (
	albumNotFoundError = errors.New("album not found")
	albumServiceInternalError = errors.New("album-service 500 error")
)
