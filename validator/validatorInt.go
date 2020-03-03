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

type ParamInt struct {
	Min     int
	Max     int
	Useable []int
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func ValidateIntMin(optIf interface{}, paramIf interface{}) error {
	opt := optIf.(argumentOption.Option)
	min := paramIf.(ParamInt).Min
	val, err := opt.GetValue()
	if err != nil {
		return err
	}
	if val.(int) < min {
		return errors.New(
			fmt.Sprintf(
				"Invalid value of --%v -%v %v. Value %v is smaller than min %v.",
				opt.LongKey, opt.ShortKey, val.(int),
				val.(int), min))
	}
	return nil
}

func ValidateIntMax(optIf interface{}, paramIf interface{}) error {
	opt := optIf.(argumentOption.Option)
	max := paramIf.(ParamInt).Max
	val, err := opt.GetValue()
	if err != nil {
		return err
	}
	if val.(int) > max {
		return errors.New(
			fmt.Sprintf(
				"Invalid value of --%v -%v %v. Value %v is bigger than max %v.",
				opt.LongKey, opt.ShortKey, val.(int),
				val.(int), max))
	}
	return nil
}

func ValidateInt(optIf interface{}, paramIf interface{}) error {
	if err := ValidateIntMin(optIf, paramIf); err != nil {
		return err
	}
	if err := ValidateIntMax(optIf, paramIf); err != nil {
		return err
	}
	return nil
}
