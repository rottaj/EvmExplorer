package program

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/rottaj/GoEthExplorer/ui"
)

func BuildAssemblyFromSol(filePath string) []byte {
	out, err := exec.Command("solc", "--opcodes", filePath).Output()
	if err != nil {
		log.Fatal(err)
	}

	return out

}

func getStepsFromOpcodes(contractOpcodes []string) [][]string {
	for i, op := range contractOpcodes { // Splices contstuctor operations --> Keeps only main code
		if op == "Opcodes:" {
			contractOpcodes = contractOpcodes[i+1:]
		}
	}
	var steps [][]string
	for i := 0; i < len(contractOpcodes)-1; i++ {
		var temp []string
		temp = append(temp, contractOpcodes[i])
		if strings.HasPrefix(contractOpcodes[i+1], "0x") {
			temp = append(temp, contractOpcodes[i+1])
			i++
		}
		steps = append(steps, temp)
	}

	return steps
}

func RunProgram(filePath string) {
	opcodes := BuildAssemblyFromSol(filePath)
	contractOpcodes := strings.Fields(string(opcodes))

	steps := getStepsFromOpcodes(contractOpcodes)
	fmt.Println(steps)
	app := ui.InitializeMainViewer()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
