package optionList

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mozzzzy/arguments/v2/argumentOption"
)

/*
 * Types
 */

type OptionList struct {
	options []argumentOption.Option
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Methods
 */

func (optList OptionList) findOptByKey(key string) (*argumentOption.Option, error) {
	// If key has prefix "-", remove them
	for ; strings.HasPrefix(key, "-"); key = key[1:] {
	}

	if len(key) > 1 {
		return optList.findOptByLongKey(key)
	}
	if len(key) == 1 {
		return optList.findOptByShortKey(key)
	}
	return nil, errors.New(fmt.Sprintf("Invalid key format \"%v\"", key))
}

func (optList OptionList) findOptByLongKey(longKey string) (*argumentOption.Option, error) {
	// If long key has prefix "-", remove them
	for ; strings.HasPrefix(longKey, "-"); longKey = longKey[1:] {
	}

	for index := 0; index < len(optList.options); index++ {
		if optList.options[index].LongKey == longKey {
			return &optList.options[index], nil
		}
	}
	return nil, errors.New("Specified option not found.")
}

func (optList OptionList) findOptByShortKey(shortKey string) (*argumentOption.Option, error) {
	// If short key has prefix "-", remove them
	for ; strings.HasPrefix(shortKey, "-"); shortKey = shortKey[1:] {
	}

	for index := 0; index < len(optList.options); index++ {
		if optList.options[index].ShortKey == shortKey {
			return &optList.options[index], nil
		}
	}
	return nil, errors.New("Specified option not found.")
}

/*
 * Public Methods
 */

func (optList *OptionList) AddOption(newOpt argumentOption.Option) error {
	validatedOpt, err := argumentOption.New(newOpt)
	if err != nil {
		return err
	}
	optList.options = append(optList.options, *validatedOpt)
	return nil
}

func (optList *OptionList) AddOptions(newOpts []argumentOption.Option) error {
	for index := 0; index < len(newOpts); index++ {
		if err := optList.AddOption(newOpts[index]); err != nil {
			return err
		}
	}
	return nil
}

func (optList *OptionList) Set(key string, value interface{}) error {
	optPtr, err := optList.findOptByKey(key)
	if err != nil {
		return err
	}
	if optPtr.Set {
		msg := "Duplicate definition of "
		if optPtr.LongKey != "" {
			msg += "--" + optPtr.LongKey + " "
		}
		if optPtr.ShortKey != "" {
			msg += "-" + optPtr.ShortKey
		}
		return errors.New(msg)
	}
	optPtr.Set = true
	if !optPtr.ValueRequired() || value == nil {
		return nil
	}
	return optPtr.SetValue(value)
}

func (optList OptionList) GetOpt(key string) (argumentOption.Option, error) {
	// Find key from long keys
	optPtr, err := optList.findOptByKey(key)
	if err != nil || optPtr == nil {
		var opt argumentOption.Option
		return opt, err
	}

	// This function does not return optPtr but *optPtr.
	// This prevents caller to modify original option data in optList.
	return *optPtr, err
}

func (optList OptionList) Get(key string) (interface{}, error) {
	// Find opt by long keys
	optPtr, err := optList.findOptByKey(key)
	// If specified key is not found, return error.
	if err != nil {
		return nil, err
	}
	// If requested option and its default value are not set, return error
	return optPtr.GetValue()
}

func (optList OptionList) GetInt(key string) (int, error) {
	var zeroVal int
	value, err := optList.Get(key)
	if err != nil {
		return zeroVal, err
	}
	integer, ok := value.(int)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf("Value of option \"%v\" is not int.", key))
	}
	return integer, nil
}

func (optList OptionList) GetString(key string) (string, error) {
	var zeroVal string
	value, err := optList.Get(key)
	if err != nil {
		return zeroVal, err
	}
	str, ok := value.(string)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf("Value of option \"%v\" is not string.", key))
	}
	return str, nil
}

func (optList OptionList) IsSet(key string) bool {
	opt, err := optList.findOptByKey(key)
	// If requested key is not found, return false.
	if err != nil {
		return false
	}
	return opt.Set
}

func (optList OptionList) Validate() error {
	for _, opt := range optList.options {
		if err := opt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This function returns "--<long key> -<short key> <value type>"
func (optList OptionList) GetOptKeys() []string {
	keys := []string{}
	for _, opt := range optList.options {
		key := ""
		// long key
		if opt.LongKey != "" {
			key += "--" + opt.LongKey
		}
		// short key
		if opt.ShortKey != "" {
			key += " "
			key += "-" + opt.ShortKey
		}
		// value place holder
		if opt.ValueType != "" {
			key += " "
			key += opt.ValueType
		}
		if key == "" {
			continue
		}
		keys = append(keys, key)
	}
	return keys
}

func (optList OptionList) String() string {
	str := ""
	if len(optList.options) == 0 {
		return str
	}

	str += "  Options\n"
	indent := "    "

	var optStrs []string
	for _, opt := range optList.options {
		optStrs = append(optStrs, opt.String())
	}
	maxStrLen := getMaxStrLen(optStrs)

	for index, opt := range optList.options {
		str += indent
		// key and value
		str += optStrs[index]
		// space between value and description
		for i := 0; i < maxStrLen-len(optStrs[index]); i++ {
			str += " "
		}
		// description
		if opt.Description != "" {
			str += " : "
			str += opt.Description
		}
		// required
		if opt.Required {
			str += " (required)"
		}
		// default value
		if opt.ValueType == "" || opt.DefaultValue == nil {
			str += "\n"
			continue
		}
		switch opt.ValueType {
		case "string":
			defaultValueStr, ok := opt.DefaultValue.(string)
			if ok {
				str += fmt.Sprintf(" (default: \"%v\")", defaultValueStr)
			}
		case "int":
			defaultValueInt, ok := opt.DefaultValue.(int)
			if ok {
				str += fmt.Sprintf(" (default: %v)", defaultValueInt)
			}
		}
		str += "\n"
	}
	return str
}

/*
 * Package Private Functions
 */

func isLongOptKey(str string) bool {
	// str is short key or operand
	if !strings.HasPrefix(str, "--") {
		return false
	}
	// str is only "--"
	if len(str) == 2 {
		return false
	}
	if !regexp.MustCompile(`[a-zA-Z0-9]`).Match([]byte(str[2:])) {
		return false
	}
	return true
}

func isShortOptKey(str string) bool {
	// str is long key
	if strings.HasPrefix(str, "--") {
		return false
	}
	// str is operand
	if !strings.HasPrefix(str, "-") {
		return false
	}
	// str is not hyphen and one character
	if len(str) != 2 {
		return false
	}
	if !regexp.MustCompile(`[a-zA-Z0-9]`).Match([]byte(str[1:])) {
		return false
	}
	return true
}

func getMaxStrLen(strs []string) int {
	maxLen := 0
	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}
	return maxLen
}

/*
 * Public Functions
 */

func IsOptKey(str string) bool {
	return isLongOptKey(str) || isShortOptKey(str)
}
