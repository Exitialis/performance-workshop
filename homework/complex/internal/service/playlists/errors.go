package playlists

import "github.com/pkg/errors"

var (
	playlistNotFoundError = errors.New("playlist not found")
	playlistServiceInternalError = errors.New("playlist-service 500 error")
)
