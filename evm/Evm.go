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

// Still need uint256 (32 bytes) padding. Because go supports up to 64 bits,
// split padding into 4 slices of 8 bytes? []int64[]int64[]int64[]int64 --> big.Int
// Use go-ethereum/uint256?

func (evm *Evm) Debug(step int) {
	evm.Stack = nil
	for i := 0; i <= step-1; i++ {
		currentStep := evm.Steps[i]
		opCode := opcodes.StringToOpcode[currentStep.Mnemonic]
		// Stack Instructions
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

		// Arithmetic Instructions
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
		if opCode == opcodes.ADDMOD {
			evm.addMod()
		}
		if opCode == opcodes.MULMOD {
			evm.mulMod()
		}
		if opCode == opcodes.EXP {
			evm.exp()
		}

		// Comparison Instructions
		if opCode == opcodes.EQ {
			evm.equal()
		}
		if opCode == opcodes.ISZERO {
			evm.isZero()
		}
		if opCode == opcodes.LT || opCode == opcodes.SLT {
			evm.lessThan()
		}
		if opCode == opcodes.GT || opCode == opcodes.SGT {
			evm.greaterThan()
		}
		// Control Instructions
		if opCode == opcodes.STOP {
			evm.stop()
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
	evm.Gas += opcodes.MOD.StaticGas // update
	evm.Pc += 1
}

func (evm *Evm) addMod() {
	x := evm.pop()
	y := evm.pop()
	z := evm.pop()
	sum := big.NewInt(0)
	if z == big.NewInt(0) {
		return
	} else {
		sum = (sum.Add(x, y))
		sum = sum.Mod(sum, z)
	}
	evm.Stack = append(evm.Stack, sum)
	evm.Pc += 1
}

func (evm *Evm) mulMod() {
	x := evm.pop()
	y := evm.pop()
	z := evm.pop()
	sum := big.NewInt(0)
	if z.Int64() == 0 {
		sum = big.NewInt(0)
	} else {
		sum = (sum.Mul(x, y))
		sum = sum.Mod(sum, z)
	}
	evm.Stack = append(evm.Stack, sum)
	evm.Pc += 1
}

func (evm *Evm) exp() {
	base := evm.pop()
	exponent := evm.pop()
	_ = base
	_ = exponent
	evm.Pc += 1
}

func (evm *Evm) lessThan() {
	left := evm.pop()
	right := evm.pop()
	res := big.NewInt(0)
	if left.Int64() < right.Int64() {
		res = big.NewInt(1)
	}
	evm.Stack = append(evm.Stack, res)
	evm.Pc += 1

}

func (evm *Evm) greaterThan() {
	left := evm.pop()
	right := evm.pop()
	res := big.NewInt(0)
	if left.Int64() > right.Int64() {
		res = big.NewInt(1)
	}
	evm.Stack = append(evm.Stack, res)
	evm.Pc += 1
}

func (evm *Evm) equal() {
	left := evm.pop()
	right := evm.pop()
	res := big.NewInt(0)
	if left.Int64() == right.Int64() {
		res = big.NewInt(1)
	}
	evm.Stack = append(evm.Stack, res)
	evm.Pc += 1
}

func (evm *Evm) isZero() {
	x := evm.pop()
	res := big.NewInt(0)
	if x == res {
		res = big.NewInt(1)
	}
	evm.Stack = append(evm.Stack, res)
	evm.Pc += 1
}

// Needs implementing

func (evm *Evm) blockHash() {
	block_number := evm.pop()
	_ = block_number
	//if len(evn)
}

func (evm *Evm) coinBase() {
	evm.Pc += 1
}
