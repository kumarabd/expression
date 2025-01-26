package main

import (
	"fmt"

	"github.com/kumarabd/expression/evaluator"
)

func main() {
	engine := evaluator.New()

	infixExpression := "( 3 + 4 ) * 2"
	result, err := engine.InfixEvaluate(infixExpression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result: ", result)

	infixExpression = "( a + b ) + c"
	result, err = engine.InfixEvaluate(infixExpression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result: ", result)

	infixExpression = "( 1 == 1 ) && ( c == c )"
	err = engine.Validate(infixExpression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	result, err = engine.InfixEvaluate(infixExpression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result: ", result)
}
