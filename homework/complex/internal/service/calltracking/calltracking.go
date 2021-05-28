package calltracking

import (
	"context"
	"github.com/exitialis/workshop/homework/complex/internal/clients/calltracking"

	"github.com/pkg/errors"

)

type Service struct {
	client Client
}

func NewCalltrackingService(client Client) *Service {
	return &Service{
		client: client,
	}
}

func (c *Service) Check(userID int64, realPhone string) bool {
	if userID % 5 == 0 {
		return true
	}

	return false
}

func (c *Service) GetPhone(ctx context.Context, req GetPhoneIn) (string, error) {
	request := calltracking.GetPhoneIn{
		Phone:     req.RealPhone,
		AccountID: req.UserID,
		ItemID:    req.ItemID,
	}

	response, err := c.client.GetVirtual(request)

	if err != nil {
		return "", errors.Wrap(GetPhoneError{
			ItemID: req.ItemID,
			UserID: req.UserID,
		}, err.Error())
	}

	if response == nil {
		return "", GetPhoneError{
			ItemID: req.ItemID,
			UserID: req.UserID,
		}
	}

	if response.Phone.VirtualPhone == "" {
		return "", PhoneNotFound{
			ItemID: req.ItemID,
			UserID: req.UserID,
		}
	}

	return response.Phone.VirtualPhone, nil
}
