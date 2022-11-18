package program

import (
	"log"
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
func getStepsFromOpcodes(contractOpcodes []string) [][]string {
	var steps [][]string
	//var gas []int
	//gas[0] = 21000 // Initialize gas price at 21000
	for i := 0; i < len(contractOpcodes)-1; i++ {
		_, isOpcode := opcodes.StringToOpcode[contractOpcodes[i]]
		if isOpcode {
			var temp []string
			temp = append(temp, contractOpcodes[i])
			if strings.HasPrefix(contractOpcodes[i+1], "0x") {
				temp = append(temp, contractOpcodes[i+1])
				i++
			}
			steps = append(steps, temp)

		}
	}
	return steps
}

func RunProgram(filePath string) {

	var evm evm.Evm

	contractOpcodes := BuildAssemblyFromSol(filePath) // Get opcodes

	opcodeSteps := getStepsFromOpcodes(contractOpcodes)
	evm.Ops = opcodeSteps
	evm.Pc = 0
	evm.Debug(3) // Debug program w/ Step

	// InitializeMainViewer arguments: opcodeSteps, stack, memory, PC?
	// Program computes everything then passes to frontend.
	// When user wants to update viewer, rerender with spliced data ( handled in ui w/ key.Pressed )

	app := ui.InitializeMainViewer(&evm)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
