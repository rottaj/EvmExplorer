package evm

import (
	"math/big"
)

func (evm *Evm) pop() {
	evm.Stack = evm.Stack[:len(evm.Stack)-1]
	evm.Pc += 1
}

func (evm *Evm) push(data *big.Int, num_bytes int) {
	evm.Stack = append(evm.Stack, data)
	evm.Pc += (1 + num_bytes)
}
