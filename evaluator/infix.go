package evaluator

import (
	"fmt"
	"strings"
)

// InfixEvaluate processes the infix expression and returns the result
func (pe *Evaluator) InfixEvaluate(infix string) (Operand, error) {
	expression, err := pe.InfixToPostfix(infix)
	if err != nil {
		return 0, err
	}
	return pe.PostfixEvaluate(expression)
}

// InfixToPostfix converts an infix expression to postfix
func (pe *Evaluator) InfixToPostfix(infix string) (string, error) {
	operatorStack := make([]string, 0)
	output := make([]string, 0)
	tokens := strings.Fields(infix)

	for _, token := range tokens {
		operation, ok := OperationList[token]
		if !ok {
			switch token {
			case "(":
				operatorStack = append(operatorStack, token)
			case ")":
				for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
					output = append(output, operatorStack[len(operatorStack)-1])
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
				if len(operatorStack) == 0 {
					return "", fmt.Errorf("mismatched parentheses")
				}
				operatorStack = operatorStack[:len(operatorStack)-1] // Pop "("
			default:
				output = append(output, token)
			}
			continue
		}
		for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" &&
			OperationList[operatorStack[len(operatorStack)-1]].precedence >= operation.precedence {
			output = append(output, operatorStack[len(operatorStack)-1])
			operatorStack = operatorStack[:len(operatorStack)-1]
		}
		operatorStack = append(operatorStack, token)
	}

	for len(operatorStack) > 0 {
		if operatorStack[len(operatorStack)-1] == "(" {
			return "", fmt.Errorf("mismatched parentheses")
		}
		output = append(output, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return strings.Join(output, " "), nil
}

func (pe *Evaluator) Validate(s string) error {
	tokens := strings.Fields(s)
	if len(tokens) == 0 {
		return fmt.Errorf("empty expression")
	}

	for i, token := range tokens {
		if _, isOp := OperationList[token]; isOp {
			if i == 0 || i == len(tokens)-1 {
				return fmt.Errorf("operator '%s' at invalid position", token)
			}
		} else if token == "(" || token == ")" {
			continue
		} else {
			// Check if token is a valid number or string
			if len(strings.Split(token, "(")) > 1 || len(strings.Split(token, ")")) > 1 {
				return fmt.Errorf("invalid token: %s, add spaces appropriately", token)
			}
		}
	}

	// Check for proper spacing
	if !strings.Contains(s, " ") || strings.Contains(s, "  ") {
		return fmt.Errorf("improper spacing between tokens")
	}

	return nil
}
