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
		if opCode == opcodes.ADD {
			evm.add()
		}
		if opCode == opcodes.SUB {
			evm.sub()
		}
		if opCode == opcodes.MUL {
			evm.mul()
		}
		if opCode == opcodes.DIV || opCode == opcodes.SDIV {
			evm.div()
		}
		if opCode == opcodes.MOD || opCode == opcodes.SMOD {
			evm.mod()
		}

		evm.Steps[i].Gas = evm.Gas
		evm.Steps[i].Pc = evm.Pc
	}

}

func (evm *Evm) add() {
	x := evm.pop()
	y := evm.pop()
	sum := big.NewInt(0)
	sum = sum.Add(x, y)
	evm.Stack = append(evm.Stack, sum)
	evm.Gas += opcodes.ADD.StaticGas // update
	evm.Pc += 1
}

func (evm *Evm) sub() {
	x := evm.pop()
	y := evm.pop()
	sum := big.NewInt(0)
	sum = sum.Sub(x, y)
	evm.Stack = append(evm.Stack, sum)
	evm.Gas += opcodes.ADD.StaticGas // update
	evm.Pc += 1
}

func (evm *Evm) mul() {
	x := evm.pop()
	y := evm.pop()
	sum := big.NewInt(0)
	sum = sum.Mul(x, y)
	evm.Stack = append(evm.Stack, sum)
	evm.Gas += opcodes.ADD.StaticGas // update
	evm.Pc += 1
}

func (evm *Evm) div() {
	x := evm.pop()
	y := evm.pop()
	sum := big.NewInt(0)
	sum = sum.Div(x, y)
	evm.Stack = append(evm.Stack, sum)
	evm.Gas += opcodes.ADD.StaticGas // update
	evm.Pc += 1
}

func (evm *Evm) mod() {
	x := evm.pop()
	y := evm.pop()
	sum := big.NewInt(0)
	sum = sum.Mod(x, y)
	evm.Stack = append(evm.Stack, sum)
	evm.Gas += opcodes.ADD.StaticGas // update
	evm.Pc += 1
}

func (evm *Evm) addmod() {

}
