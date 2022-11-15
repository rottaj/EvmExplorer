package ui

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/opcodes"
)

func createOperationStepsUI(operationPanel *tview.Flex, opcodeSteps [][]string) *tview.Flex {
	table := tview.NewTable().SetBorders(false)

	for i, op := range opcodeSteps {

		temp := opcodes.StringToOpcode[op[0]]
		fmt.Println(i, op, byte(temp.Op))
		table.SetCell(i, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))
		if len(op) > 1 {
			table.SetCell(i, 1, tview.NewTableCell(op[0]+" "+op[1]))
		} else {
			table.SetCell(i, 1, tview.NewTableCell(op[0]))
		}
		table.SetCell(i, 3, tview.NewTableCell(fmt.Sprintf("0x%x", temp.Op)))
	}

	operationPanel.AddItem(table, 0, 4, false)
	return operationPanel
}
