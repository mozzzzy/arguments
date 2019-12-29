package arguments

import (
	"errors"
)

func checkParamsLen(length int, validatorParams ...interface{}) error {
	if len(validatorParams) != length {
		return errors.New("Invalid length of validator's parameters.")
	}
	return nil
}
