package evm

import (
	"log"
	"math/big"
	"strconv"
	"strings"

	//"github.com/holiman/uint256"
	"github.com/rottaj/EvmExplorer/opcodes"
)

type Evm struct {
	Pc     int
	Gas    int
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
		// Stack Ops
		if opcodes.IsPush(opCode) {

			//val := strings.Split(evm.Ops[i][1], "0x")   // Delete str
			x := strings.Split(opCode.Mnemonic, "PUSH") // Grab PUSH_N Bytes
			byteLength, err := strconv.Atoi(x[len(x)-1])
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(currentStep.Data, byteLength)
			evm.push(currentStep.Data, byteLength) // pushdata to stack
		}
		if opcodes.IsDup(opCode) {
			x := strings.Split(opCode.Mnemonic, "DUP") // Grab PUSH_N Bytes
			byteLength, err := strconv.Atoi(x[len(x)-1])
			if err != nil {
				log.Fatal(err)
			}
			evm.dup(byteLength)
		}
		if opcodes.IsSwap(opCode) {
			x := strings.Split(opCode.Mnemonic, "SWAP") // Grab PUSH_N Bytes
			byteLength, err := strconv.Atoi(x[len(x)-1])
			if err != nil {
				log.Fatal(err)
			}
			evm.swap(byteLength) // passes in item number
		}
		if opCode == opcodes.POP {
			x := evm.pop()
			_ = x
		}
		if opCode == opcodes.MSTORE {
			evm.mstore()
		}
		// Comparison Ops
		if opCode == opcodes.ISZERO { // Needs testing
			x := evm.pop()
			_ = x
		}
		//if
		//if opCode
		evm.Steps[i].Gas = evm.Gas
		evm.Steps[i].Pc = evm.Pc
	}

}
