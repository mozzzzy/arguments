package arguments

import (
	"flag"
	"testing"
	"time"
)

/*
 * Variables
 */

/*
 * Functions
 */

/*
 * Tests
 */

func TestParseRequiredOrNotRequired(t *testing.T) {
	t.Run("Parse required and specified rule", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false
		specifiedValue := true
		specifiedValueStr := "true"

		required_and_specified_rule_name := "required-and-specified-" + valueType

		optionRules := []Option{
			{
				LongKey:      required_and_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
				Required:     true,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(required_and_specified_rule_name, specifiedValueStr)

		errRequiredAndSpecified := args.Parse()
		noError(t, errRequiredAndSpecified)

		actualRequiredAndSpecified := new(bool)
		args.Get(required_and_specified_rule_name, actualRequiredAndSpecified)
		match(t, specifiedValue, *actualRequiredAndSpecified)
	})

	t.Run("Parse required and not specified rule", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false

		required_and_not_specified_rule_name := "required-and-not-specified-" + valueType

		optionRules := []Option{
			{
				LongKey:      required_and_not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
				Required:     true,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		errRequiredAndNotSpecified := args.Parse()
		withError(t, errRequiredAndNotSpecified)

		actualRequiredAndNotSpecified := new(bool)
		args.Get(required_and_not_specified_rule_name, actualRequiredAndNotSpecified)
		match(t, defaultValue, *actualRequiredAndNotSpecified)
	})

	t.Run("Parse not required and specified rule", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false
		specifiedValue := true
		specifiedValueStr := "true"

		not_required_and_specified_rule_name := "not-required-and-specified-" + valueType

		optionRules := []Option{
			{
				LongKey:      not_required_and_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
				Required:     false,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(not_required_and_specified_rule_name, specifiedValueStr)

		errNotRequiredAndSpecified := args.Parse()
		noError(t, errNotRequiredAndSpecified)

		actualNotRequiredAndSpecified := new(bool)
		args.Get(not_required_and_specified_rule_name, actualNotRequiredAndSpecified)
		match(t, specifiedValue, *actualNotRequiredAndSpecified)
	})

	t.Run("Parse not required and not specified rule", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false

		not_required_and_not_specified_rule_name := "not-required-and-not-specified-" + valueType

		optionRules := []Option{
			{
				LongKey:      not_required_and_not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
				Required:     false,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		errNotRequiredAndNotSpecified := args.Parse()
		noError(t, errNotRequiredAndNotSpecified)

		actualNotRequiredAndNotSpecified := new(bool)
		args.Get(not_required_and_not_specified_rule_name, actualNotRequiredAndNotSpecified)
		match(t, defaultValue, *actualNotRequiredAndNotSpecified)
	})
}

// NOTE
// When we use Arguments in source codes (but test codes),
// we should execute `args.Parse()` method to parse argument options.
// But this test code execute `flag.Set()` instead.
// So we don't have to execute `args.Parse()` here.
func TestParsePartiallyDefinedRule(t *testing.T) {
	t.Run("Parse rule that does not have its short key", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false
		specifiedValue := true
		specifiedValueStr := "true"

		rule_name := "only-long-" + valueType

		optionRules := []Option{
			{
				LongKey:      rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
				Required:     true,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(rule_name, specifiedValueStr)

		errRequiredAndSpecified := args.Parse()
		noError(t, errRequiredAndSpecified)

		actual := new(bool)
		args.Get(rule_name, actual)
		match(t, specifiedValue, *actual)
	})

	t.Run("Parse rule that does not have its long key", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false
		specifiedValue := true
		specifiedValueStr := "true"

		rule_name := "s"

		optionRules := []Option{
			{
				ShortKey:     rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
				Required:     true,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(rule_name, specifiedValueStr)

		errRequiredAndSpecified := args.Parse()
		noError(t, errRequiredAndSpecified)

		actual := new(bool)
		args.Get(rule_name, actual)
		match(t, specifiedValue, *actual)
	})

	t.Run("Parse specified rule that does not have its default value", func(t *testing.T) {
		valueType := "bool"
		specifiedValue := true
		specifiedValueStr := "true"

		long_rule_name := "specified-without-default-value-" + valueType
		short_rule_name := "d"

		optionRules := []Option{
			{
				LongKey:     long_rule_name,
				ShortKey:    short_rule_name,
				ValueType:   valueType,
				Description: "some descriptions",
				Required:    true,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(long_rule_name, specifiedValueStr)

		errRequiredAndSpecified := args.Parse()
		noError(t, errRequiredAndSpecified)

		actual := new(bool)
		args.Get(long_rule_name, actual)
		match(t, specifiedValue, *actual)
	})

	t.Run("Parse not specified rule that does not have its default value", func(t *testing.T) {
		valueType := "bool"

		long_rule_name := "not-specified-without-default-value-" + valueType
		short_rule_name := "e"

		optionRules := []Option{
			{
				LongKey:     long_rule_name,
				ShortKey:    short_rule_name,
				ValueType:   valueType,
				Description: "some descriptions",
				Required:    false,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		errRequiredAndSpecified := args.Parse()
		noError(t, errRequiredAndSpecified)

		actual := new(bool)
		args.Get(long_rule_name, actual)
		match(t, false, *actual)
	})

	t.Run("Parse specified rule that does not have Required", func(t *testing.T) {
		valueType := "bool"
		specifiedValue := true
		specifiedValueStr := "true"

		long_rule_name := "specified-without-required-" + valueType
		short_rule_name := "r"

		optionRules := []Option{
			{
				LongKey:     long_rule_name,
				ShortKey:    short_rule_name,
				ValueType:   valueType,
				Description: "some descriptions",
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(long_rule_name, specifiedValueStr)

		errRequiredAndSpecified := args.Parse()
		noError(t, errRequiredAndSpecified)

		actual := new(bool)
		args.Get(long_rule_name, actual)
		match(t, specifiedValue, *actual)
	})

	t.Run("Parse not specified rule that does not have Required", func(t *testing.T) {
		valueType := "bool"

		long_rule_name := "not-specified-without-required-" + valueType
		short_rule_name := "q"

		optionRules := []Option{
			{
				LongKey:     long_rule_name,
				ShortKey:    short_rule_name,
				ValueType:   valueType,
				Description: "some descriptions",
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		errRequiredAndSpecified := args.Parse()
		noError(t, errRequiredAndSpecified)

		actual := new(bool)
		args.Get(long_rule_name, actual)
		match(t, false, *actual)
	})

	t.Run("Parse rule that does not have its short key and long key", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false

		optionRules := []Option{
			{
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
				Required:     true,
			},
		}

		var args Arguments
		withError(t, args.AddRules(optionRules))
	})
}

func TestParseEachType(t *testing.T) {
	t.Run("Parse bool rule", func(t *testing.T) {
		valueType := "bool"
		defaultValue := false
		specifiedValue := true
		specifiedValueStr := "true"
		actual := new(bool)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})

	t.Run("Parse duration rule", func(t *testing.T) {
		valueType := "duration"
		defaultValue := time.Duration(2)
		specifiedValue := time.Duration(123)
		specifiedValueStr := "123ns"
		actual := new(time.Duration)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})

	t.Run("Parse float64 rule", func(t *testing.T) {
		valueType := "float64"
		defaultValue := 0.0
		specifiedValue := 123.456
		specifiedValueStr := "123.456"
		actual := new(float64)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})

	t.Run("Parse int rule", func(t *testing.T) {
		valueType := "int"
		defaultValue := 0
		specifiedValue := 123
		specifiedValueStr := "123"
		actual := new(int)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})

	t.Run("Parse int64 rule", func(t *testing.T) {
		valueType := "int64"
		defaultValue := int64(0)
		specifiedValue := int64(123)
		specifiedValueStr := "123"
		actual := new(int64)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})

	t.Run("Parse string rule", func(t *testing.T) {
		valueType := "string"
		defaultValue := ""
		specifiedValue := "some-value"
		specifiedValueStr := "some-value"
		actual := new(string)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})

	t.Run("Parse uint rule", func(t *testing.T) {
		valueType := "uint"
		defaultValue := uint(0)
		specifiedValue := uint(123)
		specifiedValueStr := "123"
		actual := new(uint)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})

	t.Run("Parse uint rule64", func(t *testing.T) {
		valueType := "uint64"
		defaultValue := uint64(0)
		specifiedValue := uint64(123)
		specifiedValueStr := "123"
		actual := new(uint64)

		specified_rule_name := "specify-" + valueType
		not_specified_rule_name := "not-specify-" + valueType

		optionRules := []Option{
			{
				LongKey:      specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
			{
				LongKey:      not_specified_rule_name,
				ValueType:    valueType,
				Description:  "some descriptions",
				DefaultValue: defaultValue,
			},
		}

		var args Arguments
		args.AddRules(optionRules)

		flag.Set(specified_rule_name, specifiedValueStr)

		expected := specifiedValue
		err := args.Get(specified_rule_name, actual)

		noError(t, err)
		match(t, expected, *actual)
	})
}
