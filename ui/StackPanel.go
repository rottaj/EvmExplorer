package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func createStackPanelUI(stackPanel *tview.Flex, evm *evm.Evm) *tview.Flex {
	table := tview.NewTable().SetBorders(false)

	// Iterate through ops and build Stack

	// Check opcodes and build stack
	/*
		for i := 0; i <= len(ops)-1; i++ {
			opCode := opcodes.StringToOpcode[ops[i][0]]
			if opcodes.IsPush(opCode) {
				stack.Push(ops[i][1])
			}
		}
		for i, x := range stack.StackValue {
			table.SetCell(i, 3, tview.NewTableCell(fmt.Sprintf(strconv.Itoa(i)+" "+x)))
		}
	*/
	stackPanel.AddItem(table, 0, 4, true)
	return stackPanel
}
