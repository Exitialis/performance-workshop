package singers

import "github.com/pkg/errors"

var (
	singerNotFoundError = errors.New("singer not found")
	singerServiceInternalError = errors.New("singer-service 500 error")
)
