package phones_gateway

// ValidationError определяет пользовательскую ошибку валидации
type ValidationError struct {
	message string
}

func (e *ValidationError) Error() string {
	return e.message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		message: message,
	}
}
