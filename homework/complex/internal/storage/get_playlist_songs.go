package storage

func (s *Storage) GetPlaylistSongsById(id uint64) ([]uint64, bool) {
	playlist, ok := s.playlistSongs[id]

	return playlist, ok
}