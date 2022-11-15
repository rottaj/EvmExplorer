package evm

import (
	"github.com/rottaj/EvmExplorer/opcodes"
)

type Stack struct {
	StackValue []string
}

// Adds step to Stack
func (stack *Stack) Add(b string) {
	stack.StackValue = append(stack.StackValue, b)
}

// Check Operation Codes

func IsPush(op opcodes.Opcode) bool {
	if op == opcodes.PUSH1 {
		return true
	} else {
		return false
	}
}
