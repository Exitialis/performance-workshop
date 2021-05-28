package phones

import (
	"github.com/pkg/errors"
)

type PhoneService struct {
	phonesClient Client
}

func NewPhonesService(
	phonesClient Client,
) *PhoneService {
	return &PhoneService{
		phonesClient: phonesClient,
	}
}

func (p *PhoneService) GetPhone(in GetPhoneIn) (string, error) {
	userPhone, err := p.phonesClient.GetPhoneInfoByID(in.PhoneID)

	if err != nil {
		err = errors.Wrap(GetPhoneError(in), err.Error())
		return "", err
	}

	if len(userPhone) == 0 {
		return "", PhoneNotFound(in)
	}

	if userPhone == "" {
		return "", PhoneNotFound(in)
	}

	return userPhone, nil
}
