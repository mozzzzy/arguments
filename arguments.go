package arguments

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/mozzzzy/arguments/v2/argumentOption"
	"github.com/mozzzzy/arguments/v2/argumentOperand"
	"github.com/mozzzzy/arguments/v2/operandList"
	"github.com/mozzzzy/arguments/v2/optionList"
)

/*
 * Types
 */

type Args struct {
	Executed   string
	optionList optionList.OptionList
	operandList operandList.OperandList
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (args *Args) AddOption(opt argumentOption.Option) error {
	return args.optionList.AddOption(opt)
}

func (args *Args) AddOptions(opts []argumentOption.Option) error {
	return args.optionList.AddOptions(opts)
}

func (args Args) GetOpt(key string) (interface{}, error) {
	return args.optionList.Get(key)
}

func (args Args) GetIntOpt(key string) (int, error) {
	return args.optionList.GetInt(key)
}

func (args Args) GetStringOpt(key string) (string, error) {
	return args.optionList.GetString(key)
}

func (args Args) OptIsSet(key string) bool {
	return args.optionList.IsSet(key)
}

func (args *Args) AddOperand(ope argumentOperand.Operand) error {
	return args.operandList.AddOperand(ope)
}

func (args *Args) AddOperands(opes []argumentOperand.Operand) error {
	return args.operandList.AddOperands(opes)
}

func (args Args) GetOperand(key string) (interface{}, error) {
	return args.operandList.Get(key)
}

func (args Args) GetIntOperand(key string) (int, error) {
	return args.operandList.GetInt(key)
}

func (args Args) GetStringOperand(key string) (string, error) {
	return args.operandList.GetString(key)
}

func (args Args) OperandIsSet(key string) bool {
	return args.operandList.IsSet(key)
}

func (args *Args) Parse() error {
	operandCount := 0
	operandKeys := args.operandList.GetOpeKeys()

	// Parse arguments to executed file name, options and operands
	for index := 0; index < len(os.Args); index++ {
		argStr := os.Args[index]

		// executed file name
		if index == 0 {
			args.Executed = argStr
			continue
		}

		// option
		if optionList.IsOptKey(argStr) {
			// This opt is not a pointer.
			// So even if we modify this opt, the original opt in optionList is not modified.
			opt, err := args.optionList.GetOpt(argStr)
			if err != nil {
				return errors.New(
					fmt.Sprintf(
						"Failed to get option setting of \"%v\" from option list. %v",
						argStr,
						err.Error()))
			}
			valueStr := ""
			if opt.ValueRequired() {
				index++
				if index >= len(os.Args) || optionList.IsOptKey(os.Args[index]) {
					return errors.New(
						fmt.Sprintf("option %v requires value but is not speficied.", argStr))
				}
				valueStr = os.Args[index]
			}
			var value interface{}
			switch opt.ValueType {
			case "":
				value = nil
			case "string":
				value = valueStr
			case "int":
				valueInt, err := strconv.Atoi(valueStr)
				if err != nil {
					return err
				}
				value = valueInt
			}
			if err := args.optionList.Set(argStr, value); err != nil {
				return errors.New(
					fmt.Sprintf("Failed to set option \"%v\". %v", argStr, err.Error()))
			}
			continue
		}

		// If argStr does not have prefix "--" and "-",
		// this argStr is operand.
		if operandCount >= len(operandKeys) {
			return errors.New(fmt.Sprintf("To many operands %v", argStr))
		}

		opeKey := operandKeys[operandCount]
		operandCount++

		// This operand is not a pointer.
		// So even if we modify this operand, the original operand in operandList is not modified.
		operand, err := args.operandList.GetOpe(opeKey)
		if err != nil {
			return err
		}

		var value interface{}
		switch operand.ValueType {
		case "":
			value = nil
		case "string":
			value = argStr
		case "int":
			valueInt, err := strconv.Atoi(argStr)
			if err != nil {
				return errors.New(fmt.Sprintf(
					"Failed to parse operand %v \"%v\". %v",
					opeKey,
					argStr,
					err.Error()))
			}
			value = valueInt
		}
		if err := args.operandList.Set(opeKey, value); err != nil {
			return errors.New(
				fmt.Sprintf("Failed to set operand \"%v\". %v", argStr, err.Error()))
		}
	}
	return args.Validate()
}

func (arg Args) String() string {
	str := ""
	str += "\nUsage: \n"
	str += "  " + arg.Executed + " Options Operands"
	str += "\n"

	str += "\n"
	str += arg.optionList.String()
	str += "\n"
	str += arg.operandList.String()

	return str
}

func (arg Args) Validate() error {
	if err := arg.optionList.Validate(); err != nil {
		return err;
	}

	return arg.operandList.Validate()
}

/*
 * Package Private Functions
 */

/*
 * Public Functions
 */
