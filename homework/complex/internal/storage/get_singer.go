package storage

func (s *Storage) GetSingerById(id uint64) (Singer, bool) {
	playlist, ok := s.singers[id]

	return playlist, ok
}
