package evm

type Stack struct {
	depth      int
	stackValue []byte
}

// Adds step to Stack
func (stack *Stack) Add(b byte) {
	stack.stackValue = append(stack.stackValue, b)
	stack.depth += 1
}
