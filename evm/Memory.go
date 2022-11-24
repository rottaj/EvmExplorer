package evm

import (
	"fmt"
	"strings"

	"github.com/rottaj/EvmExplorer/opcodes"
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
	var memoryLocation []string
	_ = memoryLocation

	x := start_position.Uint64()
	x = x * 2

	//fmt.Println("Memory Location:", x)

	for i := 0; i < int(x); i++ {
		memoryLocation = append(memoryLocation, "0")
	}

	// Build data
	z := strings.Split(t1, "")

	fmt.Println(z)
	for i := 0; i <= 64-len(z); i++ {
		memoryLocation = append(memoryLocation, "0")
	}
	for i, data := range z {
		_ = i
		memoryLocation = append(memoryLocation, data)
	}

	//fmt.Println(memoryLocation, len(memoryLocation))
	// Need to update this...
	evm.Memory = memoryLocation
	evm.Gas += opcodes.MSTORE.StaticGas
	evm.Pc += 1
}
