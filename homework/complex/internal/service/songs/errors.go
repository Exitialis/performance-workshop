package songs

import "github.com/pkg/errors"

var (
	songNotFoundError = errors.New("song not found")
	songServiceInternalError = errors.New("song-service 500 error")
)
