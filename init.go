package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rottaj/GoEthExplorer/program"
)

func main() {
	args := os.Args
	fmt.Println(args[1:])
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(root)
	fmt.Println()
	filePath := root + "/" + string(args[1:][0])
	program.RunProgram(filePath)
}
