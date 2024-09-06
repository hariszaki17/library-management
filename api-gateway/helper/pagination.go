package helper

import (
	"errors"
)

func ValidatePageLimit(page, limit int) error {
	if page < 1 && limit < 1 {
		return errors.New("Invalid request, page and limit must be greater than 0")
	}
}
