package phones

type PhoneID int64

type Client interface {
	GetPhoneInfoByID(phoneId int64) (string, error)
}

type GetPhoneIn struct {
	PhoneID int64
}
