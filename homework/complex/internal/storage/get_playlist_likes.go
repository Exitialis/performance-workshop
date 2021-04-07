package storage

func (s *Storage) GetPlaylistLikes(userId uint64) ([]uint64, bool) {
	likes, ok := s.userPlaylistLikes[userId]

	return likes, ok
}
