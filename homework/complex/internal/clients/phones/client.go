package phones

import (
	"fmt"
	"math/rand"
	"time"
)

type PhonesClient struct {}

func NewPhonesClient() *PhonesClient {
	rand.Seed(time.Now().Unix())
	return &PhonesClient{}
}

func (*PhonesClient) GetPhoneInfoByID(phoneId int64) (string, error) {
	phone := rand.Intn(9_999_999_999 - 9_000_000_000) + 9_000_000_000

	return fmt.Sprintf("%d", phone), nil
}
