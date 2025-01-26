package evaluator

import (
	"fmt"
	"strings"
)

type Operand interface{}
type OperandFunction func(a, b Operand) (Operand, error)

type Operation struct {
	function   OperandFunction
	precedence int
}

var OperationList = map[string]Operation{
	"+": {
		function:   add,
		precedence: 98,
	},
	"-": {
		function:   subtract,
		precedence: 98,
	},
	"*": {
		function:   multiply,
		precedence: 99,
	},
	"/": {
		function:   divide,
		precedence: 99,
	},
	"&&": {
		function:   and,
		precedence: 97,
	},
	"||": {
		function:   or,
		precedence: 97,
	},
	"==": {
		function:   equality,
		precedence: 96,
	},
	"!=": {
		function:   notEquality,
		precedence: 96,
	},
}

func add(a, b Operand) (Operand, error) {
	switch a := a.(type) {
	case float64:
		if bInt, ok := b.(float64); ok {
			return a + bInt, nil
		}
	case string:
		if bStr, ok := b.(string); ok {
			return a + bStr, nil
		}
	}
	return nil, fmt.Errorf("invalid operands for addition")
}

func subtract(a, b Operand) (Operand, error) {
	if aInt, ok := a.(float64); ok {
		if bInt, ok := b.(float64); ok {
			return aInt - bInt, nil
		}
	}
	return nil, fmt.Errorf("invalid operands for subtraction")
}

func multiply(a, b Operand) (Operand, error) {
	if aInt, ok := a.(float64); ok {
		if bInt, ok := b.(float64); ok {
			return aInt * bInt, nil
		}
	}
	return nil, fmt.Errorf("invalid operands for multiplication")
}

func divide(a, b Operand) (Operand, error) {
	if aInt, ok := a.(float64); ok {
		if bInt, ok := b.(float64); ok {
			if bInt == 0 {
				return nil, fmt.Errorf("division by zero")
			}
			return aInt / bInt, nil
		}
	}
	return nil, fmt.Errorf("invalid operands for division")
}

func and(a, b Operand) (Operand, error) {
	switch a := a.(type) {
	case bool:
		if bBool, ok := b.(bool); ok {
			return a && bBool, nil
		}
	case float64:
		if bFloat, ok := b.(float64); ok {
			return (a != 0) && (bFloat != 0), nil
		}
	case int:
		if bFloat, ok := b.(int); ok {
			return (a != 0) && (bFloat != 0), nil
		}
	}
	return nil, fmt.Errorf("invalid operands for AND operation")
}

func or(a, b Operand) (Operand, error) {
	switch a := a.(type) {
	case bool:
		if bBool, ok := b.(bool); ok {
			return a || bBool, nil
		}
	case float64:
		if bFloat, ok := b.(float64); ok {
			return (a != 0) || (bFloat != 0), nil
		}
	case int:
		if bFloat, ok := b.(int); ok {
			return (a != 0) || (bFloat != 0), nil
		}
	}
	return nil, fmt.Errorf("invalid operands for OR operation")
}

func equality(a, b Operand) (Operand, error) {
	switch a := a.(type) {
	case string:
		return strings.Compare(a, b.(string)) == 0, nil
	case bool, float64:
		return a == b, nil
	}
	return nil, fmt.Errorf("invalid operands for equality comparison")
}

func notEquality(a, b Operand) (Operand, error) {
	switch a := a.(type) {
	case string:
		return strings.Compare(a, b.(string)) != 0, nil
	case bool, float64:
		return a != b, nil
	}
	return nil, fmt.Errorf("invalid operands for inequality comparison")
}
