package phones_gateway

import (
	"context"
	"github.com/exitialis/workshop/homework/complex/internal/service/calltracking"
	"github.com/exitialis/workshop/homework/complex/internal/service/phones"
)

const TypeBase = 1
const TypeCalltracking = 3

// Интерфейс для получения телефона
type PhonesGateway interface {
	GetPhone(context context.Context, in GetPhoneIn) (*GetPhoneOut, error)
	Validate(in GetPhoneIn) error
}

type PhonesService interface {
	GetPhone(in phones.GetPhoneIn) (string, error)
}

type CalltrackingService interface {
	Check(userID int64, phone string) bool
	GetPhone(ctx context.Context, req calltracking.GetPhoneIn) (string, error)
}

type GetPhoneIn struct {
	PhoneID         int64
	UserID          int64
	CategoryID      int64 // Category id of item
	ItemID          int64
	PhoneDisplayLoc string // Место, где будет показан номер, например: site.item.contact, messenger.chat, mobile.search, etc
}

type GetPhonesRowIn struct {
	PhoneID    int64
	UserID     int64
	CategoryID int64 // Category id of item
	ItemID     int64
}

type GetPhoneOut struct {
	Phone string
	Type  int64
}

// для bulk операций требуется дополнительно ItemID чтобы понять, от какого объявления телефон
type GetPhonesOut struct {
	GetPhoneOut
	ItemID int64
}
