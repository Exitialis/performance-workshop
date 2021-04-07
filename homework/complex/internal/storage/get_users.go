package storage

func (s *Storage) GetUserById(id uint64) (User, bool) {
	user, ok := s.users[id]
	return user, ok
}
