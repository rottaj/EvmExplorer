package evm

func (evm *Evm) Pop() {
	evm.Stack = evm.Stack[:len(evm.Stack)-1]
	evm.Pc += 1
}

func (evm *Evm) Push(data int, num_bytes int) {
	evm.Stack = append(evm.Stack, data)
	evm.Pc += (1 + num_bytes)
}
