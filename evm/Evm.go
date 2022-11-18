package evm

import (
	_"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/rottaj/EvmExplorer/opcodes"
)

type Evm struct {
	Pc     int
	Stack  []*big.Int
	Memory []byte
	Ops    [][]string // value push at init.. only used for calling.
}

func (evm *Evm) Debug(step int) {
	evm.Stack = nil
	for i := 0; i <= step-1; i++ {
		//fmt.Println(evm.Ops[i])
		opCode := opcodes.StringToOpcode[evm.Ops[i][0]]
		if opcodes.IsPush(opCode) {

			val := strings.Split(evm.Ops[i][1], "0x")   // Delete str
			x := strings.Split(opCode.Mnemonic, "PUSH") // Grab bytes
			z := val[len(val)-1]
			byteLength, err := strconv.Atoi(x[len(x)-1])
			if err != nil {
				log.Fatal(err)
			}
			data := new(big.Int)
			data.SetString(z, 16)
			evm.push(data, byteLength) // pushdata to stack
		}

		//fmt.Println(evm.Stack)

	}
	evm.Pc = step
}
