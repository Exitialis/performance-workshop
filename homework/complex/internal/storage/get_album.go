package storage

func (s *Storage) GetAlbumById(id uint64) (Album, bool) {
	playlist, ok := s.albums[id]

	return playlist, ok
}
