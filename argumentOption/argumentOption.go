package argumentOption

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

type Option struct {
	LongKey        string
	ShortKey       string
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

func validateRule(opt Option) error {
	if opt.LongKey == "" && opt.ShortKey == "" {
		return errors.New("Long key or short key is required.")
	}
	if opt.Required && opt.DefaultValue != nil {
		return errors.New(
			fmt.Sprintf(
				"Required option --%v -%v can't be specified its default value.",
				opt.LongKey, opt.ShortKey))
	}
	return nil
}

/*
 * Public Functions
 */

func New(opt Option) (*Option, error) {
	// Validate
	if err := validateRule(opt); err != nil {
		return nil, err
	}
	return &opt, nil
}

/*
 * Package Private Methods
 */

/*
 * Public Methods
 */

func (opt *Option) GetValue() (interface{}, error) {
	if !opt.Set && opt.DefaultValue == nil {
		return nil, errors.New(
			fmt.Sprintf(
				"No value and no default value for --%v -%v are set.",
				opt.LongKey, opt.ShortKey))
	}
	if !opt.Set {
		return opt.DefaultValue, nil
	}
	return opt.Value, nil
}

func (opt *Option) SetValue(value interface{}) error {
	if value == nil {
		return errors.New("nil is invalid for SetValue func's param.")
	}
	switch opt.ValueType {
	case "":
	case "string":
		str, ok := value.(string)
		if ok {
			opt.Value = str
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to option. "+
						"The ValueType is string. "+
						"But specified value is %T.", value))
		}
	case "int":
		integer, ok := value.(int)
		if ok {
			opt.Value = integer
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to option. "+
						"The ValueType is int. "+
						"But specified value is %T.", value))
		}
	}
	return nil
}

func (opt Option) ValueRequired() bool {
	return opt.ValueType == "string" || opt.ValueType == "int"
}

func (opt Option) Validate() error {
	// Required but not set
	if opt.Required && opt.Value == nil {
		return errors.New(
			fmt.Sprintf("Required option --%v -%v is not provided.", opt.LongKey, opt.ShortKey))
	}

	// Execute validator
	if opt.Validator != nil {
		return opt.Validator(opt, opt.ValidatorParam)
	}
	return nil
}

func (opt Option) String() string {
	str := ""
	// long key
	if opt.LongKey != "" {
		str += "--" + opt.LongKey
	}
	// short key
	if opt.ShortKey != "" {
		if len(str) != 0 {
			str += " "
		}
		str += "-" + opt.ShortKey
	}
	// value type
	if opt.ValueType != "" {
		str += " "
		str += opt.ValueType
	}
	return str
}
