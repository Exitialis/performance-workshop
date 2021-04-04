package handlers

import (
	"net/http"
	"regexp"
)

type ProfileHandler struct {

}

type ProfileResponse struct {
	UserID int64 `json:"userId"`
	Settings map[int64]Option `json:"settings"`
	EnableSidebar *bool `json:"enableSidebar,omitempty"` // Поле, необходимое только для ios приложения
}

type Option struct {
	Value *int64
	IntValues []int64
	RangeValue *RangeValue
}

type RangeValue struct {
	From int64
	To int64
}

func (h *ProfileHandler) Handle(w http.ResponseWriter, req *http.Request) {
	iphoneUAReg := "i(Phone|Pad|Pod)"
	ua := req.Header.Get("User-Agent")

	// Игнорируем ошибку, т.к. если не удалось узнать, что это iphone, то не важно
	isIphone, _ := regexp.Match(iphoneUAReg, []byte(ua))
	if isIphone {

	}
}
