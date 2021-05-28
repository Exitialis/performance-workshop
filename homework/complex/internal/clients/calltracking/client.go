package calltracking

import (
	"fmt"
	"math/rand"
	"time"
)

type CalltrackingClient struct {}

type Phone struct {
	VirtualPhone string
}

type CalltrackingResponse struct {
	Phone Phone
}

type GetPhoneIn struct {
	Phone string
	AccountID int64
	ItemID int64
}

func NewCalltrackingClient() *CalltrackingClient {
	rand.Seed(time.Now().Unix())
	return &CalltrackingClient{}
}

func (*CalltrackingClient) GetVirtual(in GetPhoneIn) (*CalltrackingResponse, error) {
	phone := rand.Intn(9_999_999_999 - 9_000_000_000) + 9_000_000_000

	return &CalltrackingResponse{
		Phone: Phone{VirtualPhone: fmt.Sprintf("%d", phone)},
	}, nil
}
