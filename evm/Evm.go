package evm

import (
	"fmt"
)

type Evm struct {
	Pc     int
	Stack  []int
	Memory []byte
	Ops    [][]string // value push at init.. only used for calling.
}

func (evm *Evm) PushToStack(b int) {

}

func (evm *Evm) Debug(step int) {
	for i := 0; i <= step-1; i++ {
		fmt.Println("FOO", evm.Ops[i])
	}
	evm.Pc = step
}
