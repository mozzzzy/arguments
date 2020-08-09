package operandList

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"

	"github.com/mozzzzy/arguments/argumentOperand"
)

/*
 * Types
 */

type OperandList struct {
	operands []argumentOperand.Operand
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Methods
 */

func (opeList OperandList) findOpeByKey(key string) (*argumentOperand.Operand, error) {
	for index := 0; index < len(opeList.operands); index++ {
		if opeList.operands[index].Key == key {
			return &opeList.operands[index], nil
		}
	}
	return nil, errors.New("Specified operand not found.")
}

/*
 * Public Methods
 */

func (opeList *OperandList) AddOperand(newOpe argumentOperand.Operand) error {
	validatedOpe, err := argumentOperand.New(newOpe)
	if err != nil {
		return err
	}
	opeList.operands = append(opeList.operands, *validatedOpe)
	return nil
}

func (opeList *OperandList) AddOperands(newOpes []argumentOperand.Operand) error {
	for index := 0; index < len(newOpes); index++ {
		if err := opeList.AddOperand(newOpes[index]); err != nil {
			return err
		}
	}
	return nil
}

func (opeList *OperandList) Set(key string, value interface{}) error {
	opePtr, err := opeList.findOpeByKey(key)
	if err != nil {
		return err
	}
	if opePtr.Set {
		msg := "Duplicate definition of " + opePtr.Key
		return errors.New(msg)
	}
	opePtr.Set = true
	if value == nil {
		return nil
	}
	return opePtr.SetValue(value)
}

func (opeList OperandList) GetOpe(key string) (argumentOperand.Operand, error) {
	// Find key from long keys
	opePtr, err := opeList.findOpeByKey(key)
	if err != nil || opePtr == nil {
		var ope argumentOperand.Operand
		return ope, err
	}

	// This function does not return opePtr but *opePtr.
	// This prevents caller to modify original operand data in opeList.
	return *opePtr, err
}

func (opeList OperandList) Get(key string) (interface{}, error) {
	// Find ope by keys
	opePtr, err := opeList.findOpeByKey(key)
	// If specified key is not found, return error.
	if err != nil {
		return nil, err
	}
	// If requested operand and its default value are not set, return error
	return opePtr.GetValue()
}

func (opeList OperandList) GetInt(key string) (int, error) {
	var zeroVal int
	value, err := opeList.Get(key)
	if err != nil {
		return zeroVal, err
	}
	integer, ok := value.(int)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf("Value of operand \"%v\" is not int.", key))
	}
	return integer, nil
}

func (opeList OperandList) GetString(key string) (string, error) {
	var zeroVal string
	value, err := opeList.Get(key)
	if err != nil {
		return zeroVal, err
	}
	str, ok := value.(string)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf("Value of operand \"%v\" is not string.", key))
	}
	return str, nil
}

func (opeList OperandList) IsSet(key string) bool {
	ope, err := opeList.findOpeByKey(key)
	// If requested key is not found, return false.
	if err != nil {
		return false
	}
	return ope.Set
}

func (opeList OperandList) Validate() error {
	for _, ope := range opeList.operands {
		if err := ope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This function returns "<key> <value type>"
func (opeList OperandList) GetOpeKeys() []string {
	keys := []string{}
	for _, ope := range opeList.operands {
		key := ""
		if ope.Key != "" {
			key += ope.Key
		}
		if len(key) == 0 {
			continue
		}
		keys = append(keys, key)
	}
	return keys
}

func (opeList OperandList) String() string {
	str := ""
	if len(opeList.operands) == 0 {
		return str
	}

	str = "  Operands\n"
	indent := "    "

	var opeKeys []string
	for _, operand := range opeList.operands {
		opeKeys = append(opeKeys, operand.Key)
	}

	for i, operand := range opeList.operands {
		opeKeys[i] += " (" + operand.ValueType + ")"
	}

	maxKeyLen := getMaxStrLen(opeKeys)

	for index, operand := range opeList.operands {
		str += indent
		// key and value
		str += opeKeys[index]
		// space between value and description
		for i := 0; i < maxKeyLen-len(opeKeys[index]); i++ {
			str += " "
		}

		// description
		if operand.Description != "" {
			str += " : "
			str += operand.Description
		}
		// required
		if operand.Required {
			str += " (required)"
		}
		// default value
		if operand.DefaultValue == nil {
			str += "\n"
			continue
		}
		switch operand.ValueType {
		case "string":
			defaultValueStr, ok := operand.DefaultValue.(string)
			if ok {
				str += fmt.Sprintf(" (default: \"%v\")", defaultValueStr)
			}
		case "int":
			defaultValueInt, ok := operand.DefaultValue.(int)
			if ok {
				str += fmt.Sprintf(" (default: %v)", defaultValueInt)
			}
		}
		str += "\n"
	}
	return str
}



/*
 * Private Functions
 */

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
