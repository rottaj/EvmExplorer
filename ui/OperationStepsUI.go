package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func createOperationStepsUI(operationPanel *tview.Flex, opcodes [][]string) *tview.Flex {
	table := tview.NewTable().SetBorders(true)
	/*
		cols := [3]string{"STEP", "MNEMONIC", "OPCODE"}
		curr := 0
		for r := 0; r <= len(opcodes); r++ {
			for c := 0; c < len(cols); c++ {
				table.SetCell(r, c, tview.NewTableCell(cols[curr]))
			}
			curr = (curr + 1) % len(opcodes)
		}
	*/
	for i, op := range opcodes {
		fmt.Println(i, op)
		table.SetCell(i, 0, tview.NewTableCell(fmt.Sprint(i)))
		table.SetCell(i, 1, tview.NewTableCell(op[0]))
	}

	operationPanel.AddItem(table, 0, 4, false)
	return operationPanel
}
