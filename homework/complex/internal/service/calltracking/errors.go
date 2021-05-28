package calltracking

import "fmt"

type PhoneNotFound struct {
	ItemID     int64
	UserID     int64
}

func (p PhoneNotFound) Error() string {
	return fmt.Sprintf("Номер не найден itemId: %d userId: %d",
		p.ItemID,
		p.UserID,
	)
}

type GetPhoneError struct {
	ItemID     int64
	UserID     int64
}

func (p GetPhoneError) Error() string {
	return fmt.Sprintf("Ошибка при получении анонимного номера, itemId: %d, userId: %d",
		p.ItemID,
		p.UserID,
	)
}
