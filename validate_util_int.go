package arguments

import (
	"errors"
	"fmt"
)

/*
 * Validators for int
 */

func ValidateIntMin(
	optKey string, optValue interface{}, validatorParams ...interface{}) error {

	// Check length of validatorParams
	checkParamsLenErr := checkParamsLen(1, validatorParams)
	if checkParamsLenErr != nil {
		return checkParamsLenErr
	}

	// Get min and max from validatorParams
	var validatorParamSlice []int = validatorParams[0].([]int)
	var min int = validatorParamSlice[0]

	// Check that option's value is smaller than min
	if *(optValue.(*int)) < min {
		return errors.New(
			fmt.Sprintf(
				"Value of -%v %v is smaller than minimum %v",
				optKey,
				*(optValue.(*int)),
				min,
			),
		)
	}
	return nil
}

func ValidateIntMax(
	optKey string, optValue interface{}, validatorParams ...interface{}) error {

	// Check length of validatorParams
	checkParamsLenErr := checkParamsLen(1, validatorParams)
	if checkParamsLenErr != nil {
		return checkParamsLenErr
	}

	// Get min and max from validatorParams
	var validatorParamSlice []int = validatorParams[0].([]int)
	var max int = validatorParamSlice[0]

	// Check that option's value is bigger than max
	if *(optValue.(*int)) > max {
		return errors.New(
			fmt.Sprintf(
				"Value of -%v %v is bigger than maximum %v",
				optKey,
				*(optValue.(*int)),
				max,
			),
		)
	}
	return nil
}

func ValidateIntMinMax(
	optKey string, optValue interface{}, validatorParams ...interface{}) error {

	// Check length of validatorParams
	checkParamsLenErr := checkParamsLen(1, validatorParams)
	if checkParamsLenErr != nil {
		return checkParamsLenErr
	}

	// Get min and max from validatorParams
	var validatorParamSlice []int = validatorParams[0].([]int)
	var min int = validatorParamSlice[0]
	var max int = validatorParamSlice[1]

	errValidateIntMin := ValidateIntMin(optKey, optValue, []int{min})
	if errValidateIntMin != nil {
		return errValidateIntMin
	}

	errValidateIntMax := ValidateIntMax(optKey, optValue, []int{max})
	if errValidateIntMax != nil {
		return errValidateIntMax
	}
	return nil
}
