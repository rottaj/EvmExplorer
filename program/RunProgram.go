package program

import (
	"fmt"
	"log"
	"os/exec"
)

func RunProgram(filePath string) {
	test := BuildAssemblyFromSol(filePath)
	fmt.Println("Test", string(test))
	//ui.MainViewer()
}

func BuildAssemblyFromSol(filePath string) []byte {
	out, err := exec.Command("solc", "--opcodes", filePath).Output()
	if err != nil {
		log.Fatal(err)
	}

	return out

}
