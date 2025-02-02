package error_wrapper

type errorCategory struct {
	statusCode    int
	maskedMessage string
}

func NewCategory(category int, maskedMessage string) *errorCategory {
	return &errorCategory{
		statusCode:    category,
		maskedMessage: maskedMessage,
	}
}

func (e *errorCategory) StatusCode() int {
	return e.statusCode
}

func (e *errorCategory) MaskedMessage() string {
	return e.maskedMessage
}
