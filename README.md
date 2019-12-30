# arguments
Argument option parser for go. This is object base wrapper of the standard flag package.  

## What is this
I always have to write codes for following operations.  
1. Parse argument options. (This functionality has already provided by standard flag package.)
2. Stop program if following situations.
	- Our required options are not provided.
	- The specified options are invalid to use.

So I aggregate codes to provide above operations to structures and methods.  
Please read [Usage](#Usage) to know the detail.


## Installation
We can install this package by `go get`.
```sh
$ go get github.com/mozzzzy/arguments
```

## Usage
### Import
First, we have to import the package.
```go
import "github.com/mozzzzy/arguments"
```

### Set parse rules
Second, we have to specify the parse rules of argument options before parse them.  
This is an example to parse `-help` and `-port` options.
```go
package main

import (
	"fmt"
	"github.com/mozzzzy/arguments"
)

func main() {
  // Define parse rules
	optionRules := []arguments.Option{
		{
			Key:          "help",
			ValueType:    "bool",
			Description:  "show Usage message and exit.",
			DefaultValue: false,
			Required:     false,
		},
		{
			Key:          "port",
			ValueType:    "int",
			Description:  "specify port number.",
			DefaultValue: false,
			Required:     false,
		},
	}

	var args arguments.Arguments

	// Add rules
	args.AddRules(optionRules)
```

A rule is defined by `arguments.Option` structure.  
`arguments.Option` has following public fields.  
- _Key_  
	- `string`  
	- Specify key of this option. If you want to parse `-help` option, set `"help"`.  
- _ValueType_
	- `string`
	- Specify this option's value type. `"bool"`, `"duration"`, `"float64"`, `"int"`, `"int64"`, `"string"`, `"uint"`, `"uint64"` is available.
- _Description_
	- `string`
	- Specify the description of this option.
- _DefaultValue_
	- `interface{}`
	- Default value of this option. If this option is not set, this default value is used.
- _Required_
	- `bool`
	- Specify that this option is required(`true`) or optional(`false`).
- _Validator_
	- `func(string, interface{}, interface{}) error`
	- Validator of the option value. See [Validate option's value](#Validate-option's-value) for detail.
- _ValidatorParam_
	- `interface{}`
	- Validator's parameter. See [Validate option's value](#Validate-option's-value) for detail.



## Get option value
We can get option value by `Get(key string, valuePtr *interface{}) error` function.  
Or if we can get bool value, we can get by `IsTrue(key string) bool` function.
```go
package main

import (
	"fmt"
	"github.com/mozzzzy/arguments"
)

func main() {

	...

	// Parse flag options
	parseErr := args.Parse()
	if parseErr != nil {
		fmt.Println(parseErr)
		args.Usage()
		return
	}

	// Check true/false
	if args.IsTrue("help") {
		args.Usage()
	}

	// Get value
	portPtr := new(int)
	portErr := args.Get("port", portPtr)
	fmt.Println(*portPtr, portErr)
}
```

## Validate option's value
We can validate specified option's value when execute `Parse()` method.  
To inject validation function, we use `Option.Validator` and `Option.ValidatorParam`.  
  
For example, we can execute `arguments.ValidateIntMinMax` function to validate  `-port` option's value.
```go
package main

import (
	"fmt"
	"github.com/mozzzzy/arguments"
)

func main() {
	optionRules := []arguments.Option{
		{
			Key:             "port",
			ValueType:       "int",
			Description:     "specify port number.",
			DefaultValue:    80,
			Required:        false,
			Validator:       arguments.ValidateIntMinMax,
			ValidatorParam: []int{0, 65535},
		},
	}

	var args arguments.Arguments

	// Add rules
	args.AddRules(optionRules)

	// Parse flag options
	parseErr := args.Parse()
	if parseErr != nil {
		fmt.Println(parseErr)
		return
	}

	vPtr := new(int)
	vErr := args.Get("port", vPtr)
	fmt.Println(*vPtr, vErr)
}
```
In this example, the first element of Options.ValidatorParam is the minimum of -port value. And the second element is maximum.  
If you run with `-port 65536`, `aras.Parse()` method returns error like following result.
```sh
$ go run valiate.go -port 65536
Value of -port 65536 is bigger than maximum 65535
```

## How to define validator function
Validator function is a function whose type is
```
func(string, interface{}, interface{}) error
```
The first parameter is the option key, and the second one is the value.  The third parameter is `Option.ValidatorParam`.
