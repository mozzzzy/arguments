package arguments

import (
	"errors"
	"flag"
	"time"
)

/*
 * Types
 */

type Arguments struct {
	Options []Option
}

type Option struct {
	DefaultValue    interface{}
	Description     string
	Key             string
	Required        bool
	set             bool
	Validator       func(string, interface{}, ...interface{}) error
	ValidatorParams interface{}
	Value           interface{}
	ValueType       string
}

/*
 * Functions
 */
func (args *Arguments) AddRule(optionRule Option) {
	args.Options = append(args.Options, optionRule)
}

func (args *Arguments) AddRules(optionRules []Option) error {
	for _, opt := range optionRules {
		args.Options = append(args.Options, opt)
		err := args.setFlag(&(args.Options[len(args.Options)-1]))
		if err != nil {
			return err
		}
	}
	return nil
}

func (args *Arguments) checkRequired() error {
	for index, opt := range args.Options {
		flag.Visit(func(flg *flag.Flag) {
			if flg.Name == opt.Key {
				args.Options[index].set = true
			}
		})
		if !args.Options[index].set && opt.Required {
			return errors.New("Required option -" + opt.Key + " is not set.")
		}
	}
	return nil
}

func (args *Arguments) Get(key string, valuePtr interface{}) error {
	errorMsg := "-" + key
	for _, opt := range args.Options {
		if opt.Key != key {
			continue
		}

		switch vp := valuePtr.(type) {
		case *bool:
			dataTypeStr := "bool"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*bool))
			return nil

		case *(time.Duration):
			dataTypeStr := "duration"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*(time.Duration)))
			return nil

		case *float64:
			dataTypeStr := "float64"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*float64))
			return nil

		case *int:
			dataTypeStr := "int"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*int))
			return nil

		case *int64:
			dataTypeStr := "int64"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*int64))
			return nil

		case *string:
			dataTypeStr := "string"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*string))
			return nil

		case *uint:
			dataTypeStr := "uint"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*uint))
			return nil

		case *uint64:
			dataTypeStr := "uint64"
			if opt.ValueType != dataTypeStr {
				return errors.New(
					errorMsg + " is not " + dataTypeStr + " type, " + opt.ValueType)
			}
			*vp = *(opt.Value.(*uint64))
			return nil

		default:
			return errors.New("Invalid second parameter type.")
		}
	}
	return errors.New(errorMsg + " is not found")
}

func (args *Arguments) IsTrue(key string) bool {
	for _, opt := range args.Options {
		if opt.Key == key {
			if opt.ValueType != "bool" {
				return false
			}
			return *(opt.Value.(*bool))
		}
	}
	return false
}

func (args *Arguments) Parse() error {
	// Parse flags
	flag.Parse()
	// Check each required options
	requiredCheckErr := args.checkRequired()
	if requiredCheckErr != nil {
		return requiredCheckErr
	}
	// Validate each options
	validateErr := args.validate()
	if validateErr != nil {
		return validateErr
	}
	return nil
}

func (args *Arguments) setFlag(opt *Option) error {
	switch opt.ValueType {
	case "bool":
		valuePtr := new(bool)
		opt.Value = valuePtr
		flag.BoolVar(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(bool),
			opt.Description,
		)
	case "duration":
		valuePtr := new(time.Duration)
		opt.Value = valuePtr
		flag.DurationVar(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(time.Duration),
			opt.Description,
		)
	case "float64":
		valuePtr := new(float64)
		opt.Value = valuePtr
		flag.Float64Var(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(float64),
			opt.Description,
		)
	case "int":
		valuePtr := new(int)
		opt.Value = valuePtr
		flag.IntVar(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(int),
			opt.Description,
		)
	case "int64":
		valuePtr := new(int64)
		opt.Value = valuePtr
		flag.Int64Var(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(int64),
			opt.Description,
		)
	case "string":
		valuePtr := new(string)
		opt.Value = valuePtr
		flag.StringVar(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(string),
			opt.Description,
		)
	case "uint":
		valuePtr := new(uint)
		opt.Value = valuePtr
		flag.UintVar(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(uint),
			opt.Description,
		)
	case "uint64":
		valuePtr := new(uint64)
		opt.Value = valuePtr
		flag.Uint64Var(
			valuePtr,
			opt.Key,
			opt.DefaultValue.(uint64),
			opt.Description,
		)
	default:
		return errors.New(
			"Failed to set rule for -" + opt.Key + ". " +
				"invalid data type: " + opt.ValueType)
	}
	return nil
}

func (args *Arguments) Usage() {
	flag.Usage()
}

func (args *Arguments) validate() error {
	// For each options
	for _, opt := range args.Options {
		// If the option is set, and validator function is specified
		if opt.set && opt.Validator != nil {
			// Execute validator
			err := opt.Validator(opt.Key, opt.Value, opt.ValidatorParams)
			// If validator return error, return it
			if err != nil {
				return err
			}
		}
	}
	// If all validator is executed successfully, return nil
	return nil
}
