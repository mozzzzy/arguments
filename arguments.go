package arguments

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mozzzzy/arguments/argumentOption"
)

/*
 * Types
 */

type Args struct {
	Executed string
	options  []argumentOption.Option
	Operands []string
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Package Private Functions
 */

func (args Args) findOptByLongKey(longKey string) *argumentOption.Option {
	for index := 0; index < len(args.options); index++ {
		if args.options[index].LongKey == longKey {
			return &args.options[index]
		}
	}
	return nil
}

func (args Args) findOptByShortKey(shortKey string) *argumentOption.Option {
	for index := 0; index < len(args.options); index++ {
		if args.options[index].ShortKey == shortKey {
			return &args.options[index]
		}
	}
	return nil
}

func getKeyStrs(opts []argumentOption.Option) []string {
	keys := []string{}
	for _, opt := range opts {
		key := "  "
		// long key
		if opt.LongKey != "" {
			key += "--" + opt.LongKey
			key += " "
		}
		// short key
		if opt.ShortKey != "" {
			key += "-" + opt.ShortKey
			key += " "
		}
		// value place holder
		if opt.ValueType != "" && opt.ValueType != "nil" {
			key += opt.ValueType
			key += " "
		}
		keys = append(keys, key)
	}
	return keys
}

func getMaxLen(strs []string) int {
	maxLen := 0
	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}
	return maxLen
}

func isLongKey(argStr string) bool {
	return strings.HasPrefix(argStr, "--")
}

func isShortKey(argStr string) bool {
	// argStr is long key
	if strings.HasPrefix(argStr, "--") {
		return false
	}
	// argStr is operand
	if !strings.HasPrefix(argStr, "-") {
		return false
	}
	// argStr is not hyphen and one character
	if len(argStr) != 2 {
		return false
	}
	if !regexp.MustCompile(`[a-zA-Z0-9]`).Match([]byte(argStr[1:])) {
		return false
	}
	return true
}

func isKey(argStr string) bool {
	return isLongKey(argStr) || isShortKey(argStr)
}

/*
 * Public Functions
 */

func (args *Args) AddOption(opt argumentOption.Option) error {
	validatedOpt, err := argumentOption.New(opt)
	if err != nil {
		return err
	}
	args.options = append(args.options, *validatedOpt)
	return nil
}

func (args *Args) AddOptions(opts []argumentOption.Option) error {
	for index := 0; index < len(opts); index++ {
		if err := args.AddOption(opts[index]); err != nil {
			return err
		}
	}
	return nil
}

func (args Args) Get(key string) (interface{}, error) {
	// Find key from long keys
	opt := args.findOptByLongKey(key)
	// If long key is not found, find key from short keys
	if opt == nil {
		opt = args.findOptByShortKey(key)
	}
	// If requested key is not found, return error.
	if opt == nil {
		return nil, errors.New(fmt.Sprintf("Required key \"%v\" is not found.", key))
	}
	// If requested option and its default value are not set, return error
	return opt.GetValue()
}

func (args Args) GetInt(key string) (int, error) {
	var zeroVal int
	value, err := args.Get(key)
	if err != nil {
		return zeroVal, err
	}
	integer, ok := value.(int)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf("Value of option \"%v\" is not int.", key))
	}
	return integer, nil
}

func (args Args) GetString(key string) (string, error) {
	var zeroVal string
	value, err := args.Get(key)
	if err != nil {
		return zeroVal, err
	}
	str, ok := value.(string)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf("Value of option \"%v\" is not string.", key))
	}
	return str, nil
}

func (args Args) IsSet(key string) bool {
	// Find key from long keys
	opt := args.findOptByLongKey(key)
	// If long key is not found, find key from short keys
	if opt == nil {
		opt = args.findOptByShortKey(key)
	}
	// If requested key is not found, return false.
	if opt == nil {
		return false
	}
	return opt.Set
}

func (args *Args) Parse() error {
	for index := 0; index < len(os.Args); index++ {
		argStr := os.Args[index]

		if index == 0 {
			args.Executed = argStr
			continue
		}

		var opt *argumentOption.Option
		if isLongKey(argStr) {
			opt = args.findOptByLongKey(argStr[2:])
		} else if isShortKey(argStr) {
			opt = args.findOptByShortKey(argStr[1:])
		} else {
			// If argStr does not have prefix "--" and "-",
			// this argStr is operand.
			args.Operands = append(args.Operands, argStr)
			continue
		}
		// If argStr has prefix "--" or "-"
		// and the option key is not found in args.options, return error.
		if opt == nil {
			return errors.New(fmt.Sprintf("Unknown option %v", argStr))
		}

		// If found option has already set, return error.
		if opt.Set == true {
			msg := "Duplicate definition of "
			if opt.LongKey != "" {
				msg += "--" + opt.LongKey + " "
			}
			if opt.ShortKey != "" {
				msg += "-" + opt.ShortKey
			}
			return errors.New(msg)
		}
		opt.Set = true

		// If found option require value, get it from next argStr.
		switch opt.ValueType {
		case "":
			continue
		case "nil":
			continue
		case "string":
			index++
			if len(os.Args) <= index || isKey(os.Args[index]) {
				return errors.New(
					fmt.Sprintf("option %v requires value but is not speficied.", argStr))
			}
			if err := opt.SetValue(os.Args[index]); err != nil {
				return err
			}
		case "int":
			index++
			if len(os.Args) <= index {
				return errors.New(
					fmt.Sprintf("option %v requires value but is not speficied.", argStr))
			}
			integer, err := strconv.Atoi(os.Args[index])
			if err != nil {
				return errors.New(
					fmt.Sprintf(
						"Invalid int value for %v \"%v\". %v", argStr, os.Args[index], err.Error()))
			}
			if err := opt.SetValue(integer); err != nil {
				return err
			}
		}
	}
	return args.Validate()
}

func (arg Args) String() string {
	str := ""
	str += "usage: \n"

	// Get strings whose formats are
	// "  --<long key> -<short key> <data type>"
	keys := getKeyStrs(arg.options)

	// Get max str length of keys
	maxLen := getMaxLen(keys)

	// Create usage string
	for index, opt := range arg.options {
		// key
		key := keys[index]
		str += key
		for index := 0; index < maxLen-len(key); index++ {
			str += " "
		}
		// description
		if opt.Description != "" {
			str += ": "
			str += opt.Description
		}
		// required
		if opt.Required == true {
			str += " (required)"
		}
		// default value
		if opt.ValueType != "" && opt.ValueType != "nil" {
			if opt.DefaultValue != nil {
				switch opt.ValueType {
				case "string":
					defaultValStr, ok := opt.DefaultValue.(string)
					if ok {
						str += fmt.Sprintf(" (default: \"%v\")", defaultValStr)
					}
				case "int":
					defaultValInt, ok := opt.DefaultValue.(int)
					if ok {
						str += fmt.Sprintf(" (default: %v)", defaultValInt)
					}
				}
			}
		}
		str += "\n"
	}
	return str
}

func (arg Args) Validate() error {
	for _, opt := range arg.options {
		if err := opt.Validate(); err != nil {
			return err
		}
	}
	return nil
}
