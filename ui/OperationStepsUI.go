package ui

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/rottaj/GoEvmExplorer/opcodes"
)

func createOperationStepsUI(operationPanel *tview.Flex, opcodeSteps [][]string) *tview.Flex {
	table := tview.NewTable().SetBorders(true)

	for i, op := range opcodeSteps {
		//fmt.Println(i, op)
		temp := opcodes.StringToOpcode(op[0])
		fmt.Println("TEMP", temp)
		table.SetCell(i, 0, tview.NewTableCell(fmt.Sprint(i)))
		table.SetCell(i, 1, tview.NewTableCell(op[0]))
	}

	operationPanel.AddItem(table, 0, 4, false)
	return operationPanel
}
