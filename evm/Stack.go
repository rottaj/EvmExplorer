package evm

import (
	"math/big"

	"github.com/rottaj/EvmExplorer/opcodes"
)

func (evm *Evm) pop() *big.Int {
	//fmt.Println("TESTPOP", evm.Stack[len(evm.Stack)])
	var poppedData *big.Int
	if len(evm.Stack) > 1 {
		poppedData = evm.Stack[len(evm.Stack)-1]
		evm.Stack = evm.Stack[:len(evm.Stack)-1]
	}
	evm.Pc += 1
	evm.Gas += opcodes.POP.StaticGas
	return poppedData
}

func (evm *Evm) push(data *big.Int, num_bytes int) {
	evm.Stack = append(evm.Stack, data)
	evm.Gas += opcodes.PUSH1.StaticGas // update
	evm.Pc += (1 + num_bytes)
}

func (evm *Evm) dup(item_number int) { // number to be duplicated to the top of stack

	if item_number < len(evm.Stack) {
		data_to_push := evm.Stack[len(evm.Stack)-item_number]
		evm.Stack = append(evm.Stack, data_to_push)

		evm.Gas += opcodes.DUP1.StaticGas
		evm.Pc += 1
	}

}

func (evm *Evm) swap(item_number int) {
	evm.Stack[0], evm.Stack[item_number] = evm.Stack[item_number], evm.Stack[0]
	evm.Gas += opcodes.SWAP1.StaticGas
	evm.Pc += 1
}
