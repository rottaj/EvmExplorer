package evm

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	//"github.com/holiman/uint256"
	"github.com/rottaj/EvmExplorer/opcodes"
)

type Evm struct {
	Pc     int
	Stack  []*big.Int
	Memory []int
	Ops    [][]string // value push at init.. only used for calling.
	Steps  []*Step
}
type Step struct {
	Pc       int
	Mnemonic string
	Data     *big.Int
	Op       byte
	Gas      int
}

func (evm *Evm) Debug(step int) {
	evm.Stack = nil
	for i := 0; i <= step-1; i++ {
		currentStep := evm.Steps[i]
		opCode := opcodes.StringToOpcode[currentStep.Mnemonic]
		if opcodes.IsPush(opCode) {

			//val := strings.Split(evm.Ops[i][1], "0x")   // Delete str
			x := strings.Split(opCode.Mnemonic, "PUSH") // Grab PUSH_N Bytes
			byteLength, err := strconv.Atoi(x[len(x)-1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(currentStep.Data, byteLength)
			evm.push(currentStep.Data, byteLength) // pushdata to stack
		}
		if opCode == opcodes.POP {
			evm.pop()
		}
		if opCode == opcodes.MSTORE {
			evm.mstore()
		}
		//fmt.Println(evm.Ops[i])
		//fmt.Println("STACK", evm.Stack)
		//operationStep.Pc = evm.Pc
		//evm.Steps = append(evm.Steps, operationStep) // to step
	}

}
