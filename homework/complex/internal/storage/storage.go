package storage

import (
	"encoding/json"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"path"
)

type Storage struct {
	songs map[uint64]Song
	singers map[uint64]Singer
	albums map[uint64]Album
	playlists map[uint64]Playlist
	users map[uint64]User

	playlistSongs map[uint64][]uint64
	userSongsLikes map[uint64][]uint64
	userAlbumLikes map[uint64][]uint64
	userPlaylistLikes map[uint64][]uint64
	userSingerLikes map[uint64][]uint64
}

type User struct {
	ID uint64 `json:"_id"`
	Username string
}

type Song struct {
	ID uint64 `json:"_id"`
	Name string
	Length uint16
	AlbumID uint64 `json:"album"`
	Artist uint64
}

type Singer struct {
	ID uint64 `json:"_id"`
	Name string
	Logo string
}

type Album struct {
	ID uint64 `json:"_id"`
	Name string
	Logo string
}

type Playlist struct {
	ID uint64 `json:"_id"`
	Name string `json:"name"`
	UserID uint64
	Logo string
}

type PlaylistSongs struct {
	SongID uint64 `json:"song_id"`
	PlaylistID uint64 `json:"playlist_id"`
}

type UserSongsLikes struct {
	UserID uint64
	SongID uint64
}

type UserAlbumLikes struct {
	UserID uint64
	AlbumID uint64
}

type UserPlaylistLikes struct {
	UserID uint64 `json:"userId"`
	PlaylistID uint64 `json:"playlist_id"`
}

type UserSingerLikes struct {
	UserID uint64
	SingerID uint64
}

func NewStorage() *Storage {
	s := Storage{
		songs:             make(map[uint64]Song),
		singers:           make(map[uint64]Singer),
		albums:            make(map[uint64]Album),
		playlists:         make(map[uint64]Playlist),
		users:             make(map[uint64]User),
		playlistSongs:     make(map[uint64][]uint64),
		userSongsLikes:    make(map[uint64][]uint64),
		userAlbumLikes:    make(map[uint64][]uint64),
		userPlaylistLikes: make(map[uint64][]uint64),
		userSingerLikes:   make(map[uint64][]uint64),
	}

	eg := errgroup.Group{}
	eg.Go(s.loadAlbums)
	eg.Go(s.loadUsers)
	eg.Go(s.loadSongs)
	eg.Go(s.loadPlaylists)
	eg.Go(s.loadSingers)
	eg.Go(s.loadUserPlaylistsLikes)
	eg.Go(s.loadPlaylistSongs)

	err := eg.Wait()
	if err != nil {
		log.Fatal(err)
	}

	return &s
}

func (s *Storage) loadUsers() error {
	var users []User
	dat, err := readData("users.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dat, &users)
	if err != nil {
		return err
	}

	for _, user := range users {
		s.users[user.ID] = user
	}

	return nil
}

func (s *Storage) loadSongs() error {
	var songs []Song
	dat, err := readData("songs.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dat, &songs)
	if err != nil {
		return err
	}

	for _, song := range songs {
		s.songs[song.ID] = song
	}

	return nil
}

func (s *Storage) loadAlbums() error {
	var albums []Album
	dat, err := readData("albums.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dat, &albums)
	if err != nil {
		return err
	}

	for _, album := range albums {
		s.albums[album.ID] = album
	}

	return nil
}

func (s *Storage) loadSingers() error {
	var singers []Singer
	dat, err := readData("singers.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dat, &singers)
	if err != nil {
		return err
	}

	for _, singer := range singers {
		s.singers[singer.ID] = singer
	}

	return nil
}

func (s *Storage) loadPlaylists() error {
	var playlists []Playlist
	dat, err := readData("playlists.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dat, &playlists)
	if err != nil {
		return err
	}

	for _, playlist := range playlists {
		s.playlists[playlist.ID] = playlist
	}

	return nil
}

func (s *Storage) loadUserPlaylistsLikes() error {
	var likes []UserPlaylistLikes
	dat, err := readData("user_playlist_likes.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dat, &likes)
	if err != nil {
		return err
	}

	for _, playlist := range likes {
		s.userPlaylistLikes[playlist.UserID] = append(s.userPlaylistLikes[playlist.UserID], playlist.PlaylistID)
	}

	return nil
}

func (s *Storage) loadPlaylistSongs() error {
	var likes []PlaylistSongs
	dat, err := readData("playlist_songs.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dat, &likes)
	if err != nil {
		return err
	}

	for _, playlist := range likes {
		s.playlistSongs[playlist.PlaylistID] = append(s.playlistSongs[playlist.PlaylistID], playlist.SongID)
	}

	return nil
}

func readData(filename string) ([]byte, error) {
	pwd, _ := os.Getwd()
	file := path.Join(pwd, "../data", filename)

	return os.ReadFile(file)
}