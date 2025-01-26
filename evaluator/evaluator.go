package evaluator

// Evaluator struct to encapsulate the evaluation logic
type Evaluator struct {
	stack []Operand
}

// NewEvaluator creates a new Evaluator instance
func New() *Evaluator {
	return &Evaluator{
		stack: make([]Operand, 0),
	}
}
