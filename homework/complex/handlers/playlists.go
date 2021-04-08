package handlers

import (
	"encoding/json"
	"github.com/exitialis/workshop/homework/complex/internal/storage"
	"golang.org/x/sync/errgroup"
	"net/http"
	"regexp"
	"strconv"
	"go.octolab.org/pointer"
)

// Ручка возвращает понравившиеся плейлисты пользователю
// и в цикле "случайно" перебирает плейлисты и ходит в еще какой-то сервис?
type FavoritePlaylistsHandler struct {
	albumClient AlbumClient
	playlistClient PlaylistClient
	singersClient SingersClient
	songsClient SongsClient

	storage PlaylistSongsStorage
}


func New(client AlbumClient, playlistClient PlaylistClient, singersClient SingersClient, songsClient SongsClient, storage PlaylistSongsStorage) *FavoritePlaylistsHandler {
	return &FavoritePlaylistsHandler{
		albumClient:    client,
		playlistClient: playlistClient,
		singersClient:  singersClient,
		songsClient:    songsClient,
		storage:        storage,
	}
}

type FavoritePlaylistsListResponse struct {
	UserID uint64 `json:"userId"`
	Playlists []Playlist
	EnableSidebar *bool `json:"enableSidebar,omitempty"` // Поле, необходимое только для ios приложения
}

type Playlist struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"` // Ссылка на файл с картинкой альбома
	Songs []Song `json:"songs"` // id песен
}

type Song struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Length uint16  `json:"length"` // длина трека в секундах
	Album Album `json:"album"`
	Artist Artist `json:"artist"`
}

type Album struct {
	Name string `json:"name"`
	Logo string `json:"logo"` // Ссылка на файл с картинкой альбома
}

type Artist struct {
	Name string
	Logo string `json:"logo"` // Ссылка на файл с картинкой исполнителя
}

type ErrorResponse struct {
	Code int64 `json:"code"`
	Message string `json:"message"`
}

func (h *FavoritePlaylistsHandler) Handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	iphoneUAReg := "i(Phone|Pad|Pod)"
	ua := req.Header.Get("User-Agent")

	// Игнорируем ошибку, т.к. если не удалось узнать, что это iphone, то не важно
	isIphone, _ := regexp.Match(iphoneUAReg, []byte(ua))

	// Предположим, что наш сервис получает запросы только после прохождения слоя api-gateway
	// и наличие заголовка X-UserId означает, что пользователь авторизован и можно считать, что это он послал запрос
	userIdStr := req.Header.Get("X-UserId")
	if userIdStr == "" {
		unauthorized(w)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		unauthorized(w)
		return
	}

	// Теперь надо сходить в сервисы:
	// 1. Плейлистов, он вернет нам список плейлистов
	// 2. Для каждого плейлиста, надо получить список песен по их Id, конечно же в цикле
	// 3. Для каждой песни, необходимо получить
	playlistIds, ok := h.storage.GetPlaylistLikes(userId)
	if !ok {
		notFound(w)
		return
	}

	resp := FavoritePlaylistsListResponse{
		UserID:        userId,
		Playlists:     []Playlist{},
	}
	if len(playlistIds) == 0 {
		err = successResponse(resp, isIphone, w)
		if err != nil {
			errResponse(500, err.Error(), w)
		}
	}

	for _, playlistId := range playlistIds {
		playlist, err := h.playlistClient.GetPlaylistById(playlistId)
		if err != nil {
			errResponse(500, err.Error(), w)
			return
		}

		songsId, ok := h.storage.GetPlaylistSongsById(playlistId)
		if !ok {
			continue
		}

		playlistToResp := Playlist{
			ID:    playlistId,
			Name: playlist.Name,
			Logo:  playlist.Logo,
		}
		for _, songId := range songsId {
			song, err := h.songsClient.GetSongById(songId)
			if err != nil {
				continue
			}

			// Не забываем про параллельность запросов
			eg := errgroup.Group{}
			var album *storage.Album
			var artist *storage.Singer
			eg.Go(func() error {
				var err error
				album, err = h.albumClient.GetAlbumById(song.AlbumID)
				if err != nil {
					return err
				}

				return nil
			})
			eg.Go(func() error {
				var err error
				artist, err = h.singersClient.GetSingerById(song.Artist)
				if err != nil {
					return err
				}

				return nil
			})

			err = eg.Wait()
			if err != nil {
				errResponse(500, err.Error(), w)
				return
			}

			playlistToResp.Songs = append(playlistToResp.Songs, Song{
				ID:     song.ID,
				Name:   song.Name,
				Length: song.Length,
				Album:  Album{
					Name:  album.Name,
					Logo:  album.Logo,
				},
				Artist: Artist{
					Name:  artist.Name,
					Logo:  artist.Logo,
				},
			})
		}

		resp.Playlists = append(resp.Playlists, playlistToResp)
	}

	err = successResponse(resp, isIphone, w)
	if err != nil {
		errResponse(500, err.Error(), w)
	}
}

func successResponse(resp FavoritePlaylistsListResponse, isIos bool, w http.ResponseWriter) error {
	if isIos {
		resp.EnableSidebar = pointer.ToBool(true)
	}

	answer, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = w.Write(answer)

	return err
}


func errResponse(code int64, message string, w http.ResponseWriter) {
	resp := ErrorResponse{
		Code:    code,
		Message: message,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(`{"code": 500, "message": "server error"}`))
		return
	}

	w.Write(res)
	return
}

func unauthorized(w http.ResponseWriter) {
	errResponse(401, "Unathorized.", w)
}

func notFound(w http.ResponseWriter) {
	errResponse(404, "Not found", w)
}
