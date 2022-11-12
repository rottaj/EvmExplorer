package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func InitializeMainViewer(ops []string) {
	box := tview.NewBox().SetBorder(true).SetTitle("Go Eth Explorer")
	opsTable := tview.NewTable()
	opsTable.SetTitle("Operations").SetBorder(true)
	for i, op := range ops {
		fmt.Println(op)
		opsTable.InsertRow(i)
	}
	tview.NewApplication().SetRoot(box, true).Run()

}

func InitilizeOperationTable(ops []string) {

}
