package arguments_test

import (
	"os"
	"testing"

	"github.com/mozzzzy/arguments/v2"
	"github.com/mozzzzy/arguments/v2/argumentOperand"
	"github.com/mozzzzy/arguments/v2/argumentOption"
)

/*
 * Variables
 */

/*
 * Functions
 */

func Match(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
}

func WithError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Should got error but nil")
	}
}

/*
 * Tests
 */

func TestValueType(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				ShortKey:  "i",
				ValueType: "int",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program", "-i", "10"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		val, getIntErr := args.GetIntOpt("i")
		Match(t, 10, val)
		NoError(t, getIntErr)
	})

	t.Run("string", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				ShortKey:  "s",
				ValueType: "string",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program", "-s", "some-string"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		val, getStringErr := args.GetStringOpt("s")
		Match(t, "some-string", val)
		NoError(t, getStringErr)
	})
}

func TestOptIsSet(t *testing.T) {
	t.Run("ShortKey", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				ShortKey: "s",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program", "-s"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		Match(t, true, args.OptIsSet("s"))
	})

	t.Run("LongKey", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				LongKey: "long",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program", "--long"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		Match(t, true, args.OptIsSet("long"))
	})

	t.Run("Not set", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				LongKey: "long",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		Match(t, false, args.OptIsSet("long"))
	})
}

func TestParse(t *testing.T) {
	t.Run("ShortKey", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				ShortKey:  "s",
				ValueType: "int",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program", "-s", "10"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		val, getIntErr := args.GetIntOpt("s")
		Match(t, 10, val)
		NoError(t, getIntErr)
	})

	t.Run("LongKey", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				LongKey:   "long",
				ValueType: "int",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program", "--long", "10"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		val, getIntErr := args.GetIntOpt("long")
		Match(t, 10, val)
		NoError(t, getIntErr)
	})

	t.Run("Use DefaultValue", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				LongKey:      "long",
				ValueType:    "int",
				DefaultValue: 10,
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		val, getIntErr := args.GetIntOpt("long")
		Match(t, 10, val)
		NoError(t, getIntErr)
	})

	t.Run("OverWrite DefaultValue", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				LongKey:      "long",
				ValueType:    "int",
				DefaultValue: 5,
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program", "--long", "10"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		val, getIntErr := args.GetIntOpt("long")
		Match(t, 10, val)
		NoError(t, getIntErr)
	})

	t.Run("Required", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				LongKey:   "long",
				ValueType: "int",
				Required:  true,
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		NoError(t, addOptErr)

		os.Args = []string{"some-program"}
		parseErr := args.Parse()
		WithError(t, parseErr)
	})
}

func TestOperand(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var args arguments.Args

		opes := []argumentOperand.Operand{
			{
				Key:       "operand1",
				ValueType: "string",
			},
		}
		addOpeErr := args.AddOperands(opes)
		NoError(t, addOpeErr)

		os.Args = []string{"some-program", "string"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		Match(t, "some-program", args.Executed)

		ope1, getStrErr1 := args.GetStringOperand("operand1")
		Match(t, "string", ope1)
		NoError(t, getStrErr1)
	})

	t.Run("int", func(t *testing.T) {
		var args arguments.Args

		opes := []argumentOperand.Operand{
			{
				Key:       "operand1",
				ValueType: "int",
			},
		}
		addOpeErr := args.AddOperands(opes)
		NoError(t, addOpeErr)

		os.Args = []string{"some-program", "10"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		Match(t, "some-program", args.Executed)

		ope1, getStrErr1 := args.GetIntOperand("operand1")
		Match(t, 10, ope1)
		NoError(t, getStrErr1)
	})

	t.Run("Use DefaultValue", func(t *testing.T) {
		var args arguments.Args

		opes := []argumentOperand.Operand{
			{
				Key:          "operand1",
				ValueType:    "int",
				DefaultValue: 10,
			},
		}
		addOpeErr := args.AddOperands(opes)
		NoError(t, addOpeErr)

		os.Args = []string{"some-program"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		ope1, getIntErr1 := args.GetIntOperand("operand1")
		Match(t, 10, ope1)
		NoError(t, getIntErr1)
	})

	t.Run("OverWrite DefaultValue", func(t *testing.T) {
		var args arguments.Args

		opes := []argumentOperand.Operand{
			{
				Key:          "operand1",
				ValueType:    "int",
				DefaultValue: 10,
			},
		}
		addOpeErr := args.AddOperands(opes)
		NoError(t, addOpeErr)

		os.Args = []string{"some-program", "20"}
		parseErr := args.Parse()
		NoError(t, parseErr)

		ope1, getIntErr1 := args.GetIntOperand("operand1")
		Match(t, 20, ope1)
		NoError(t, getIntErr1)
	})

	t.Run("Required", func(t *testing.T) {
		var args arguments.Args

		opes := []argumentOperand.Operand{
			{
				Key:       "operand1",
				ValueType: "int",
				Required:  true,
			},
		}
		addOpeErr := args.AddOperands(opes)
		NoError(t, addOpeErr)

		os.Args = []string{"some-program"}
		parseErr := args.Parse()
		WithError(t, parseErr)
	})
}
