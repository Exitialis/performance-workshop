package calltracking

import (
	"github.com/exitialis/workshop/homework/complex/internal/clients/calltracking"
)

type Client interface {
	GetVirtual(request calltracking.GetPhoneIn) (*calltracking.CalltrackingResponse, error)
}

type GetPhoneIn struct {
	RealPhone  string
	ItemID     int64
	UserID     int64
}
