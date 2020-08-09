package main

/*
 * Module Dependencies
 */

import (
	"fmt"

	"github.com/mozzzzy/arguments"
	"github.com/mozzzzy/arguments/argumentOption"
	"github.com/mozzzzy/arguments/argumentOperand"
	"github.com/mozzzzy/arguments/validator"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func main() {
	var args arguments.Args

	opt1 := argumentOption.Option{
		LongKey:        "string",
		ShortKey:       "s",
		ValueType:      "string",
		Description:    "some option.",
		Required:       true,
		Validator:      validator.ValidateString,
		ValidatorParam: validator.ParamString{Min: 3, Max: 5},
	}
	opt2 := argumentOption.Option{
		LongKey:        "int",
		ShortKey:       "i",
		ValueType:      "int",
		Description:    "some option.",
		DefaultValue:   80,
		Validator:      validator.ValidateInt,
		ValidatorParam: validator.ParamInt{Min: 10, Max: 100},
	}
	opt3 := argumentOption.Option{
		LongKey:        "bool",
		ShortKey:       "b",
		Description:    "some option.",
	}

	if err := args.AddOption(opt1); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := args.AddOptions([]argumentOption.Option {opt2, opt3}); err != nil {
		fmt.Println(err.Error())
		return
	}

	ope1 := argumentOperand.Operand{
		Key: "operand1",
		ValueType: "string",
		Description: "some operand",
		Required: true,
	}
	ope2 := argumentOperand.Operand{
		Key: "operand2",
		ValueType: "int",
		Description: "some operand",
	}
	ope3 := argumentOperand.Operand{
		Key: "operand3",
		ValueType: "int",
		Description: "some operand",
	}

	if err := args.AddOperand(ope1); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := args.AddOperands([]argumentOperand.Operand {ope2, ope3}); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := args.Parse(); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
		return
	}

	if str1, err := args.GetStringOpt("string"); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
	} else {
		fmt.Println(str1)
	}

	if integer1, err := args.GetIntOpt("int"); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
	} else {
		fmt.Println(integer1)
	}

	if args.OptIsSet("b") {
		fmt.Println("--bool -b option is set.")
	} else {
		fmt.Println("--bool -b option is not set.")
	}

	if str2, err := args.GetStringOperand("operand1"); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
		return
	} else {
		fmt.Println(str2)
	}

	if integer2, err := args.GetIntOperand("operand2"); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
		return
	} else {
		fmt.Println(integer2)
	}

	if integer3, err := args.GetIntOperand("operand3"); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
		return
	} else {
		fmt.Println(integer3)
	}
}
