package awsbilling

import (
	"errors"
	"strings"
)

// asError returns a single error with all passed or nil if passed slice is empty.
func asError(errorList []error) error {
	if len(errorList) > 0 {
		errorMessages := []string{}
		for _, err := range errorList {
			errorMessages = append(errorMessages, err.Error())
		}
		return errors.New(strings.Join(errorMessages, "\n"))
	}
	return nil
}
