package evm

import (
	"fmt"
)

func (evm *Evm) mstore() {

	start_position := evm.Stack[len(evm.Stack)-1]
	evm.Stack = evm.Stack[:len(evm.Stack)-1]
	data := evm.Stack[len(evm.Stack)-1]
	evm.Stack = evm.Stack[:len(evm.Stack)-1]
	t := fmt.Sprintf("%X", start_position)
	t1 := fmt.Sprintf("%X", data)
	_ = t
	_ = t1
	//fmt.Println("T", t, t1)
	//fmt.Println(start_position, data)
	//evm.Memory = append(evm.Memory, data)
	// write to memory
	evm.Pc += 1
}
