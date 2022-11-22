package evm

import (
	"fmt"
	"log"
	"strconv"
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
	fmt.Println("Memory Position 0x", t, "Data", t1)
	x, err := strconv.ParseInt(t, 16, 16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("T", x/2)
	var memoryLocation []int

	_ = memoryLocation
	res := strings.Split(t, "")
	for i := 0; i <= len(res)-1; i++ {
		fmt.Println("I", res[i])
	}
	// 1.) Fill x w/ 0' padding (32 bytes hex)
	// 2.) Expand memory based on position

	//fmt.Println("T", t, t1)
	//fmt.Println(start_position, data)
	//evm.Memory = append(evm.Memory, data)
	// write to memory
	evm.Gas += opcodes.MSTORE.StaticGas
	evm.Pc += 1
}
