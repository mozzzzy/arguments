# arguments
Argument option parser for go.  

## What is this
I always have to write codes for following operations.  
1. Parse argument options.
2. Handle following situations.
	- Our required options are not provided.
	- The specified options are invalid to use.

So I aggregate codes to provide above operations to this package.  

## Install
We can get this package by `go get` command.
```sh
$ go get github.com/mozzzzy/argumets
```


## Usage
We can see an example code in [examples](https://github.com/mozzzzy/arguments/tree/readme/examples) directory.
### Import
First, we have to import `arguments` and `arguments/option` like this.
```go
import (
  "github.com/mozzzzy/arguments"
  "github.com/mozzzzy/arguments/option"
)
```

### Add an option rule
To parse command line options, we have to create `arguments.Args`
 and then add an option rule by `AddOption()` method.
```go
var args arguments.Args

opt := option.Option{
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
`LongKey` and `ShortKey` is keys of the option. For example, above code add a rule for `--long-key` and `-s`.  
`int` and `string` are avairable for `ValueType`.  
`Description` is the description of the option. This is used in usage message.  
  
If we have to parse multiple options, we can use `AddOptions()` method instead.
```go
opt1 := option.Option{
	LongKey:        "string",
	ShortKey:       "s",
	ValueType:      "string",
	Description:    "some option.",
}
opt2 := option.Option{
	LongKey:        "int",
	ShortKey:       "i",
	ValueType:      "int",
	Description:    "some option.",
}

if err := args.AddOptions([]option.Option{opt1, opt2}); err != nil {
	fmt.Println(err.Error())
	return
}
```

## Parse options
After adding option rules, we can parse command line options using `Parse()` method.
```go
if err := args.Parse(); err != nil {
	fmt.Println(err.Error())
	return
}
```

## Print Usage
`arguments.Args` has `String()` method.  
We can print usage message just by printing `arguments.Args`.
```go
fmt.Println(args)
```

## Get option's value
To get value of an option, we use `GetInt()` and `GetString()` method.  
The parameter is the long key or short key.
```go
if val, err := args.GetString("long-key"); err != nil {
	fmt.Println(err.Error())
	fmt.Println(args)
} else {
	fmt.Println(val)
}
```

## Detailed usage
### Required option
We can specify required options by adding `Required` field to `option.Option`.
If the program is executed without the required option,
 `Parse()` method would fail with error.
```go
opt := option.Option{
	LongKey:        "long-key",
	ShortKey:       "s",
	ValueType:      "int",
	Description:    "some option.",
	Required:       true,
}
```

### Default value
We can also specify default value of an option.
To specify default value, we add `DefaultValue` field to `option.Option`.
```go
opt := option.Option{
	LongKey:        "long-key",
	ShortKey:       "s",
	ValueType:      "int",
	Description:    "some option.",
	DefaultValue:   80,
}
```

### Validator
We often have to validate values of specified options.  
This package can validate them easily.  
We can specify validator function and its parameter to `Validator` and `ValidatorParam`.
```go
opt := option.Option{
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
The first parameter is the `option.Option` data. The second parameter is `ValidatorParam`.
