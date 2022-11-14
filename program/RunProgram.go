package program

import (
	"log"
	"os/exec"
	"strings"

	"github.com/rottaj/GoEthExplorer/ui"
)

func RunProgram(filePath string) {
	asmRet := BuildAssemblyFromSol(filePath)
	x := strings.Fields(string(asmRet))
	for i, op := range x { // Splices contstuctor operations --> Keeps only main code
		if op == "Opcodes:" {
			x = x[i+1:]
		}
	}

	ui.InitializeMainViewer(x) // Passes OpCodes
}

func BuildAssemblyFromSol(filePath string) []byte {
	out, err := exec.Command("solc", "--opcodes", filePath).Output()
	if err != nil {
		log.Fatal(err)
	}

	return out

}
