package evm

type Stack struct {
	StackValue []string
}

// Adds step to Stack
func (stack *Stack) Push(b string) {
	stack.StackValue = append(stack.StackValue, b)
}
