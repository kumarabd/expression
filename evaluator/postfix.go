package evaluator

import (
	"fmt"
	"strconv"
	"strings"
)

// PostfixEvaluate processes the postfix expression and returns the result
func (pe *Evaluator) PostfixEvaluate(expression string) (Operand, error) {
	pe.stack = pe.stack[:0] // Clear the stack
	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if err := pe.processToken(token); err != nil {
			return 0, err
		}
	}

	if len(pe.stack) != 1 {
		return 0, fmt.Errorf("invalid expression: too many operands")
	}

	return pe.stack[0], nil
}

// processToken handles individual tokens in the expression
func (pe *Evaluator) processToken(token string) error {
	token = strings.TrimSpace(token)
	operandFunction, ok := OperationList[token]
	if !ok {
		return pe.pushOperand(token)
	}
	return pe.performPostfixOperation(operandFunction.function)
}

// pushOperand converts a string token to a float and pushes it to the stack if its a number else it will push the token
func (pe *Evaluator) pushOperand(token string) error {
	num, err := strconv.ParseFloat(token, 64)
	if err != nil {
		pe.stack = append(pe.stack, token)
		return nil
	}
	pe.stack = append(pe.stack, num)
	return nil
}

// performOperation executes the arithmetic operation
func (pe *Evaluator) performPostfixOperation(function OperandFunction) error {
	if len(pe.stack) < 2 {
		return fmt.Errorf("invalid expression: not enough operands")
	}

	b, a := pe.stack[len(pe.stack)-1], pe.stack[len(pe.stack)-2]
	pe.stack = pe.stack[:len(pe.stack)-2]

	result, err := function(a, b)
	if err != nil {
		return err
	}
	pe.stack = append(pe.stack, result)
	return nil
}
