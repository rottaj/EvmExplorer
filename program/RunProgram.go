package program

import (
	"log"
	"math/big"
	"os/exec"
	"strings"

	"github.com/rottaj/EvmExplorer/evm"
	"github.com/rottaj/EvmExplorer/opcodes"
	"github.com/rottaj/EvmExplorer/ui"
)

func BuildAssemblyFromSol(filePath string) []string {
	out, err := exec.Command("solc", "--opcodes", filePath).Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(out))

}

// Takes in string from BuildAssemblyFromSol
// Returns OpcodeSteps, PC, Gas,
func getStepsFromOpcodes(contractOpcodes []string) []*evm.Step {

	var steps []*evm.Step
	for i := 0; i < len(contractOpcodes)-1; i++ {
		_, isOpcode := opcodes.StringToOpcode[contractOpcodes[i]]
		if isOpcode {
			opCode := opcodes.StringToOpcode[contractOpcodes[i]]
			var operationStep = &evm.Step{
				Mnemonic: opCode.Mnemonic,
				Gas:      opCode.StaticGas,
				Op:       opCode.Op,
				//Pc: int, // Updated in debug
			}
			if strings.HasPrefix(contractOpcodes[i+1], "0x") {
				val := strings.Split(contractOpcodes[i+1], "0x") // Delete str
				data := new(big.Int)
				data.SetString(val[len(val)-1], 16)

				//fmt.Println("BAR", val[len(val)-1])
				//fmt.Println("BAR", data)
				operationStep.Data = data
				i++
			}
			//fmt.Println("FOOO", operationStep)
			steps = append(steps, operationStep)
		}

	}
	return steps
}

func RunProgram(filePath string) {

	var evm evm.Evm

	// Program computes steps then passes to frontend. (*evm compiles opcode)
	contractOpcodes := BuildAssemblyFromSol(filePath)   // Get opcodes
	opcodeSteps := getStepsFromOpcodes(contractOpcodes) // Get steps from opcodes

	evm.Steps = opcodeSteps
	evm.Pc = 0
	evm.Gas = 21000
	evm.Debug(len(evm.Steps))
	//evm.Debug(5)
	app := ui.InitializeMainViewer(&evm)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
