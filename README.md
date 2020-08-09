# arguments
Argument option parser for go.  

## What is this
I always had to write codes for following operations.  
1. Parse argument options.
2. Handle following situations.
	- Our required options are not provided.
	- The specified options are invalid to use.

So I aggregate codes to provide above operations to this package.  

## Install
We can get this package by `go get` command.
```sh
$ go get github.com/mozzzzy/arguments/v2
```
Or if we use `go.mod`, we can write like following snippet.
```sh
module <your module name>

go 1.13

require(
  github.com/mozzzzy/arguments/v2 v2.0.1
)
```
And then execute `go get`.
```sh
$ go get
```

## Usage
We can see an example code in [examples](https://github.com/mozzzzy/arguments/tree/master/examples) directory.

### Import
First, we have to import some modules.
```go
import (
  "github.com/mozzzzy/arguments/v2"
  "github.com/mozzzzy/arguments/v2/argumentOption"  // if you parse options (like -h --help)
  "github.com/mozzzzy/arguments/v2/argumentOperand" // if you parse operands
)
```

### Handle options (like --opt VALUE or -o VALUE)
#### Add an option rule
To parse command line options, we have to create `arguments.Args`
 and then add an option rule by `AddOption()` method.
```go
var args arguments.Args

opt := argumentOption.Option{
	LongKey:        "long-key",
	ShortKey:       "s",
	ValueType:      "string",
	Description:    "some option.",
}
if err := args.AddOption(opt); err != nil {
	fmt.Println(err.Error())
	return
}
```
##### LongKey and ShortKey
`LongKey` and `ShortKey` is keys of the option.  
For example, above code add a rule for `--long-key` and `-s`.  

##### ValueType
`int` and `string` are available for `ValueType`.  
  
If we have to parse option that doesn't have value (like --help),  
we should skip configuring `ValueType` field.
```go
var args arguments.Args

opt := argumentOption.Option{
	LongKey:        "help",
	ShortKey:       "h",
	Description:    "show help message and exit.",
}
if err := args.AddOption(opt); err != nil {
	fmt.Println(err.Error())
	return
}
```

##### Description
`Description` is the description of the option. This is used in usage message.

##### Required
`Required: true` specifies the option is required.  
If required option is not set, `args.Parse()` method returns error.

##### DefaultValue
`DefaultValue` specifies the default value of the option.  
If the option is not specified, the default value is used.

##### Validator and ValidatorParam
We often have to validate option values.  
We can validate them easily.  
```go
opt := argumentOption.Option{
	LongKey:        "long-key",
	ShortKey:       "s",
	ValueType:      "int",
	Description:    "some option.",
	DefaultValue:   80,
	Validator:      validator.ValidateInt,
	ValidatorParam: validator.ParamInt{Min: 0, Max: 100},
}
```
Validator function should be `func (interface{}, interface{}) error`.  
The first parameter is the `argumentOption.Option` data. The second parameter is `ValidatorParam`.

#### Add multiple option rules at once
we can add multiple option rules at once by `AddOptions()` method.
```go
opt1 := argumentOption.Option{
	LongKey:        "string",
	ShortKey:       "s",
	ValueType:      "string",
	Description:    "some option.",
}
opt2 := argumentOption.Option{
	LongKey:        "int",
	ShortKey:       "i",
	ValueType:      "int",
	Description:    "some option.",
}

if err := args.AddOptions([]argumentOption.Option{opt1, opt2}); err != nil {
	fmt.Println(err.Error())
	return
}
```

#### Parse options
After adding option rules, we can parse command line options using `Parse()` method.
```go
var args arguments.Args

opt := argumentOption.Option{
	LongKey:        "long-key",
	ShortKey:       "s",
	ValueType:      "string",
	Description:    "some option.",
}
if err := args.AddOption(opt); err != nil {
	fmt.Println(err.Error())
	return
}

if err := args.Parse(); err != nil {
	fmt.Println(err.Error())
	return
}
```

#### Get option's value
To get value of parsed options, we use `GetIntOpt()` `GetStringOpt()` and `GetOpt()` method.  
The parameter is the long key or short key.
```go
// GetStringOpt()
if val, err := args.GetStringOpt("long-key"); err != nil {
	fmt.Println(err.Error())
	fmt.Println(args)
} else {
	fmt.Println(val)
}

// GetIntOpt()
if val, err := args.GetIntOpt("long-key"); err != nil {
	fmt.Println(err.Error())
	fmt.Println(args)
} else {
	fmt.Println(val)
}

// GetOpt()
if val, err := args.GetOpt("long-key"); err != nil {
	fmt.Println(err.Error())
	fmt.Println(args)
} else {
	valInt, ok := val.(int)
	if ok {
		fmt.Println(valInt)
	} else {
		fmt.Println("value of --long-key is not integer.")
	}
}
```

#### Check option is set
To check an option is set, we use `OptIsSet()` method.
```
if args.OptIsSet("--help") {
	fmt.Println("--help -h option is specified.")
}
```

#### Print Usage
`arguments.Args` has `String()` method.  
We can print usage message just by printing `arguments.Args`.
```go
fmt.Println(args)
```

### Handle Operands
We can handle operands by almost the same way with options.  
The differences are following points.  
* Use `argumentOperand.Operand` instead of `argumentOption.Option`
* In `argumentOperand.Operand`, use `Key` instead of `LongKey` and `ShortKey`
* Add operand rule using `AddOperand()` and `AddOperands()` methods.
* Get operand value using `GetStringOperand()` `GetIntOperand()` and `GetOperand()`
* To check an operand is set, use `OperandIsSet()` instead of `OptionIsSet()`
  
Following code is an example.
```go
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

if str2, err := args.GetStringOperand("operand1"); err != nil {
	fmt.Println(err.Error())
	fmt.Println(args)
	return
} else {
	fmt.Println(str2)
}

if integer2, err := args.GetIntOperand("operand2"); err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(integer2)
}

if ope3, err := args.GetOperand("operand3"); err != nil {
	fmt.Println(err.Error())
} else {
	integer3, ok := ope3.(int)
	if ok {
		fmt.Println(integer3)
	} else {
		fmt.Println("value of operand3 is not integer.")
	}
}
```
