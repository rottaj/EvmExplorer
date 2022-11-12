package evm

type Stack struct {
	stackValue []byte
}

func (stack *Stack) Add(b byte) {
	stack.stackValue = append(stack.stackValue, b)
}
