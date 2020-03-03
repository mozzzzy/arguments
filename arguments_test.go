package arguments_test

import (
	"os"
	"testing"

	"github.com/mozzzzy/arguments"
	"github.com/mozzzzy/arguments/argumentOption"
	"github.com/mozzzzy/testUtil"
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
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program", "-i", "10"}
		parseErr := args.Parse()
		testUtil.NoError(t, parseErr)

		val, getIntErr := args.GetInt("i")
		testUtil.Match(t, 10, val)
		testUtil.NoError(t, getIntErr)
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
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program", "-s", "some-string"}
		parseErr := args.Parse()
		testUtil.NoError(t, parseErr)

		val, getStringErr := args.GetString("s")
		testUtil.Match(t, "some-string", val)
		testUtil.NoError(t, getStringErr)
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
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program", "-s", "10"}
		parseErr := args.Parse()
		testUtil.NoError(t, parseErr)

		val, getIntErr := args.GetInt("s")
		testUtil.Match(t, 10, val)
		testUtil.NoError(t, getIntErr)
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
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program", "--long", "10"}
		parseErr := args.Parse()
		testUtil.NoError(t, parseErr)

		val, getIntErr := args.GetInt("long")
		testUtil.Match(t, 10, val)
		testUtil.NoError(t, getIntErr)
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
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program"}
		parseErr := args.Parse()
		testUtil.NoError(t, parseErr)

		val, getIntErr := args.GetInt("long")
		testUtil.Match(t, 10, val)
		testUtil.NoError(t, getIntErr)
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
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program", "--long", "10"}
		parseErr := args.Parse()
		testUtil.NoError(t, parseErr)

		val, getIntErr := args.GetInt("long")
		testUtil.Match(t, 10, val)
		testUtil.NoError(t, getIntErr)
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
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program"}
		parseErr := args.Parse()
		testUtil.WithError(t, parseErr)
	})

}

func TestOperand(t *testing.T) {
	t.Run("Operand", func(t *testing.T) {
		opts := []argumentOption.Option{
			{
				LongKey:   "long",
				ValueType: "int",
			},
		}
		var args arguments.Args
		addOptErr := args.AddOptions(opts)
		testUtil.NoError(t, addOptErr)

		os.Args = []string{"some-program", "--long", "10", "operand1", "operand2"}
		parseErr := args.Parse()
		testUtil.NoError(t, parseErr)

		val, getIntErr := args.GetInt("long")
		testUtil.Match(t, 10, val)
		testUtil.NoError(t, getIntErr)

		testUtil.Match(t, "some-program", args.Executed)

		testUtil.Match(t, 2, len(args.Operands))
		testUtil.Match(t, "operand1", args.Operands[0])
		testUtil.Match(t, "operand2", args.Operands[1])
	})
}
