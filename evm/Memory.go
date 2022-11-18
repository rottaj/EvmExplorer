package evm

func (evm *Evm) mstore() {
	//start_position := evm.Stack[len(evm.Stack)-1]
	evm.Stack = evm.Stack[:len(evm.Stack)-1]

	// write to memory
	evm.Pc += 1
}
