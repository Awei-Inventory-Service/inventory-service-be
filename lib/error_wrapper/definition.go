package error_wrapper

import "fmt"

type ErrorDefinition struct {
	code  int
	error string

	isMasked bool
	category *errorCategory
}

func NewDefinition(code int, err string, isMasked bool, errCategory *errorCategory) *ErrorDefinition {
	return &ErrorDefinition{
		code:     code,
		error:    err,
		isMasked: isMasked,
		category: errCategory,
	}
}

func (e *ErrorDefinition) Error(err []interface{}) string {
	if e.isMasked {
		return e.category.maskedMessage
	}
	return fmt.Sprintf(e.error, err...)
}

func (e *ErrorDefinition) ActualError(err []interface{}) string {
	return fmt.Sprintf(e.error, err...)
}

func (e *ErrorDefinition) Code() int {
	return e.code
}

func (e *ErrorDefinition) IsMasked() bool {
	return e.isMasked
}
