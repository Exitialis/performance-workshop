package storage

func (s *Storage) GetSongById(id uint64) (Song, bool) {
	playlist, ok := s.songs[id]

	return playlist, ok
}
