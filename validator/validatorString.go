package validator

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"

	"github.com/mozzzzy/arguments/argumentOption"
)

/*
 * Types
 */

type ParamString struct {
	Min     int
	Max     int
	Useable []string
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func ValidateStrlenMin(optIf interface{}, paramIf interface{}) error {
	opt := optIf.(argumentOption.Option)
	min := paramIf.(ParamString).Min
	val, err := opt.GetValue()
	if err != nil {
		return err
	}
	if len(val.(string)) < min {
		return errors.New(
			fmt.Sprintf(
				"Invalid value of --%v -%v \"%v\". String length %v is shorter than min %v.",
				opt.LongKey, opt.ShortKey, val.(string),
				len(val.(string)), min))
	}
	return nil
}

func ValidateStrlenMax(optIf interface{}, paramIf interface{}) error {
	opt := optIf.(argumentOption.Option)
	max := paramIf.(ParamString).Max
	val, err := opt.GetValue()
	if err != nil {
		return err
	}
	if len(val.(string)) > max {
		return errors.New(
			fmt.Sprintf(
				"Invalid value of --%v -%v \"%v\". String length %v is longer than max %v.",
				opt.LongKey, opt.ShortKey, val.(string),
				len(val.(string)), max))
	}
	return nil
}

func ValidateString(optIf interface{}, paramIf interface{}) error {
	if err := ValidateStrlenMin(optIf, paramIf); err != nil {
		return err
	}
	if err := ValidateStrlenMax(optIf, paramIf); err != nil {
		return err
	}
	return nil
}
