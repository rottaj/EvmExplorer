package program

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

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

func getStepsFromOpcodes(contractOpcodes []string) [][]string {
	var steps [][]string
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
	contractOpcodes := BuildAssemblyFromSol(filePath) // Binary
	fmt.Println(contractOpcodes, len(contractOpcodes))
	opcodeSteps := getStepsFromOpcodes(contractOpcodes)
	_ = opcodeSteps
	fmt.Println(opcodeSteps)
	app := ui.InitializeMainViewer(opcodeSteps)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
