package main

/*
 * Module Dependencies
 */

import (
	"fmt"

	"github.com/mozzzzy/arguments"
	"github.com/mozzzzy/arguments/argumentOption"
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
	if err := args.AddOption(opt1); err != nil {
		fmt.Println(err.Error())
		return
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
	if err := args.AddOption(opt2); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := args.Parse(); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
		return
	}

	if str, err := args.GetString("string"); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
	} else {
		fmt.Println(str)
	}

	if str, err := args.GetInt("int"); err != nil {
		fmt.Println(err.Error())
		fmt.Println(args)
	} else {
		fmt.Println(str)
	}
}
