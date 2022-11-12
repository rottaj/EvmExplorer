package program

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/rottaj/GoEthExplorer/ui"
)

func RunProgram(filePath string) {
	asmRet := BuildAssemblyFromSol(filePath)
	x := strings.Fields(string(asmRet))
	for i, op := range x {
		if op == "Opcodes:" {
			x = x[i+1:]
		}
	}

	fmt.Println("x", x, len(x))
	//ui.MainViewer()
	//ui.InitilizeOperationTable(x)

	ui.InitializeMainViewer(x)
}

func BuildAssemblyFromSol(filePath string) []byte {
	out, err := exec.Command("solc", "--opcodes", filePath).Output()
	if err != nil {
		log.Fatal(err)
	}

	return out

}
