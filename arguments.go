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
	DefaultValue   interface{}
	Description    string
	LongKey        string
	Required       bool
	set            bool
	ShortKey       string
	Validator      func(string, interface{}, interface{}) error
	ValidatorParam interface{}
	Value          interface{}
	ValueType      string
}

/*
 * Constants
 */

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
			if flg.Name == opt.LongKey || flg.Name == opt.ShortKey {
				args.Options[index].set = true
			}
		})
		if !args.Options[index].set && opt.Required {
			return errors.New(
				"Required option" +
					" -" + opt.LongKey +
					" (-" + opt.ShortKey +
					") is not set.")
		}
	}
	return nil
}

func (args *Arguments) Get(key string, valuePtr interface{}) error {
	errorMsg := "-" + key
	for _, opt := range args.Options {
		if opt.LongKey != key && opt.ShortKey != key {
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
		if opt.LongKey == key || opt.ShortKey == key {
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
	// Either LongKey or ShortKey is required.
	if opt.LongKey == "" && opt.ShortKey == "" {
		return errors.New(
			"Failed to set flag. Either of long option or short option is required.")
	}

	switch opt.ValueType {
	case "bool":
		// Check DefaultValue is set or not
		var defaultValue bool
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(bool)
		}
		valuePtr := new(bool)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.BoolVar(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.BoolVar(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	case "duration":
		// Check DefaultValue is set or not
		var defaultValue time.Duration
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(time.Duration)
		}
		valuePtr := new(time.Duration)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.DurationVar(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.DurationVar(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	case "float64":
		// Check DefaultValue is set or not
		var defaultValue float64
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(float64)
		}
		valuePtr := new(float64)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.Float64Var(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.Float64Var(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	case "int":
		// Check DefaultValue is set or not
		var defaultValue int
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(int)
		}
		valuePtr := new(int)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.IntVar(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.IntVar(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	case "int64":
		// Check DefaultValue is set or not
		var defaultValue int64
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(int64)
		}
		valuePtr := new(int64)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.Int64Var(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.Int64Var(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	case "string":
		// Check DefaultValue is set or not
		var defaultValue string
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(string)
		}
		valuePtr := new(string)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.StringVar(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.StringVar(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	case "uint":
		// Check DefaultValue is set or not
		var defaultValue uint
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(uint)
		}
		valuePtr := new(uint)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.UintVar(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.UintVar(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	case "uint64":
		// Check DefaultValue is set or not
		var defaultValue uint64
		if opt.DefaultValue != nil {
			defaultValue = opt.DefaultValue.(uint64)
		}
		valuePtr := new(uint64)
		opt.Value = valuePtr
		if opt.LongKey != "" {
			flag.Uint64Var(
				valuePtr,
				opt.LongKey,
				defaultValue,
				opt.Description,
			)
		}
		if opt.ShortKey != "" {
			flag.Uint64Var(
				valuePtr,
				opt.ShortKey,
				defaultValue,
				opt.Description,
			)
		}
	default:
		return errors.New(
			"Failed to set rule for" +
				" -" + opt.LongKey +
				"(-" + opt.ShortKey +
				"). Invalid data type: " + opt.ValueType)
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
			var key string
			if opt.LongKey != "" {
				key = opt.LongKey
			} else if opt.ShortKey != "" {
				key = opt.ShortKey
			}
			err := opt.Validator(key, opt.Value, opt.ValidatorParam)
			// If validator return error, return it
			if err != nil {
				return err
			}
		}
	}
	// If all validator is executed successfully, return nil
	return nil
}
