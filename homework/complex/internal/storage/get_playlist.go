package storage

func (s *Storage) GetPlaylistById(id uint64) (Playlist, bool) {
	playlist, ok := s.playlists[id]

	return playlist, ok
}
