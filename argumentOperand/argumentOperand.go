package argumentOperand

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"
)

/*
 * Types
 */

type Operand struct {
	Key            string
	Description    string
	ValueType      string
	DefaultValue   interface{}
	Value          interface{}
	Required       bool
	Set            bool
	Validator      func(interface{}, interface{}) error
	ValidatorParam interface{}
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Package Private Functions
 */

func validateRule(ope Operand) error {
	if ope.Key == "" {
		return errors.New("Key is required.")
	}
	if ope.ValueType == "" {
		return errors.New("Value type is required.")
	}
	if ope.Required && ope.DefaultValue != nil {
		return errors.New(
			fmt.Sprintf(
				"Required operand %v can't be specified its default value.",
				ope.Key))
	}
	return nil
}

/*
 * Public Functions
 */

func New(ope Operand) (*Operand, error) {
	// Validate
	if err := validateRule(ope); err != nil {
		return nil, err
	}
	return &ope, nil
}

/*
 * Package Private Methods
 */

/*
 * Public Methods
 */

func (ope *Operand) GetValue() (interface{}, error) {
	if !ope.Set && ope.DefaultValue == nil {
		return nil, errors.New(
			fmt.Sprintf(
				"No value and no default value for %v are set.",
				ope.Key))
	}
	if !ope.Set {
		return ope.DefaultValue, nil
	}
	return ope.Value, nil
}

func (ope *Operand) SetValue(value interface{}) error {
	if value == nil {
		return errors.New("nil is invalid for SetValue func's param.")
	}
	switch ope.ValueType {
	case "string":
		str, ok := value.(string)
		if ok {
			ope.Value = str
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to operand. "+
						"The ValueType is string. "+
						"But specified value is %T.", value))
		}
	case "int":
		integer, ok := value.(int)
		if ok {
			ope.Value = integer
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to operand. "+
						"The ValueType is int. "+
						"But specified value is %T.", value))
		}
	}
	return nil
}

func (ope Operand) Validate() error {
	// Required but not set
	if ope.Required && ope.Value == nil {
		return errors.New(
			fmt.Sprintf("Required operand %v is not provided.", ope.Key))
	}

	// Execute validator
	if ope.Validator != nil {
		return ope.Validator(ope, ope.ValidatorParam)
	}
	return nil
}
