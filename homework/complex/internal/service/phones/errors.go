package phones

import "fmt"

type PhoneNotFound struct {
	PhoneID int64
}

func (p PhoneNotFound) Error() string {
	return fmt.Sprintf("Не удалось получить телефон, phoneId: %d", p.PhoneID)
}

type GetPhoneError struct {
	PhoneID int64
}

func (p GetPhoneError) Error() string {
	return fmt.Sprintf("Ошибка при получении номера телефона, phoneId: %d", p.PhoneID)
}

type GetPhonesError struct {
	PhoneIDs []int64
}

func (p GetPhonesError) Error() string {
	return fmt.Sprintf("Ошибка при получении номера телефона, phoneIds: %v", p.PhoneIDs)
}

type InvalidPhone struct {
	PhoneID int64
	Phone   string
}

func (p InvalidPhone) Error() string {
	return fmt.Sprintf("Получен некорректный номер телефона: %s, phoneId: %d", p.Phone, p.PhoneID)
}
