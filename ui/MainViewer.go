package ui

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/rottaj/GoEthExplorer/opcodes"
)

func InitializeMainViewer(ops []string) {
	box := tview.NewBox().SetBorder(true).SetTitle("Go Eth Explorer")
	opsTable := tview.NewTable()
	opsTable.SetTitle("Operations").SetBorder(true)

	var stack [][]string
	_ = stack
	for i := 0; i < len(ops)-1; i++ {
		var temp []string
		temp = append(temp, ops[i])
		if strings.HasPrefix(ops[i+1], "0x") {
			temp = append(temp, ops[i+1])
			i++
		}
		fmt.Println(temp)
		stack = append(stack, temp)
	}
	fmt.Println(stack)

	t := opcodes.StringToOpcode["STOP"]
	fmt.Println(t)
	tview.NewApplication().SetRoot(box, true).Run()

}

func InitilizeOperationTable(ops []string) {

}
