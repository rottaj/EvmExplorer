package evm

import (
	"fmt"
	"log"
	"strconv"

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
	x, err := strconv.ParseInt(t, 16, 16)
	if err != nil {
		log.Fatal(err)
	}
	var memoryLocation []string

	_ = memoryLocation
	if x*2 < 32 {
		//fmt.Println("Memory Location: 0x00", x*2)
		//fmt.Println(t, t1)
		//fmt.Println("--------------------------------------------------------------")
	} else {
		//fmt.Println("Greater than 0x00", x*2)
		//fmt.Println(t, t1)
		var temp string
		for i := len(memoryLocation); i <= int(x*2); i++ {
			temp = fmt.Sprintf(temp + "0")
			if i%32 == 0 {
				memoryLocation = append(memoryLocation, string(temp))
				//fmt.Println("TEST", temp)
				temp = ""
			}
		}
		//fmt.Println(memoryLocation)
		//fmt.Println("--------------------------------------------------------------")
	}
	/*
		res := strings.Split(t, "")
		for i := 0; i <= len(res)-1; i++ {
			fmt.Println("I", res[i])
		}
	*/
	// 1.) Fill x w/ 0' padding (32 bytes hex)
	// 2.) Expand memory based on position

	//fmt.Println("T", t, t1)
	//fmt.Println(start_position, data)
	//evm.Memory = append(evm.Memory, data)
	// write to memory
	evm.Memory = memoryLocation
	evm.Gas += opcodes.MSTORE.StaticGas
	evm.Pc += 1
}
