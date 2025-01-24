package utils

import (
	"errors"
	"regexp"
)

func IsUsernameValid(username string) error {
	// Email regex pattern
	emailRegex := `^[^@\s]+@[^@\s]+\.[^@\s]+$`

	re := regexp.MustCompile(emailRegex)

	if re.MatchString(username) {
		return errors.New("username cannot be in email format")
	}

	return nil
}
