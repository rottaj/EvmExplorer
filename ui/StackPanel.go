package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func (mainUi *MainUi) createStackPanel(evm *evm.Evm) {
	// opcodePanel is adds Table w/ opcodePanelAddItem in UI creator
	//stackPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	//stackPanel.SetBorder(true).SetTitle("Stack").SetTitleAlign(0)
	stackPanel := tview.NewTable().SetBorder(true).SetTitle("Stack")

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
	//stackPanel.AddItem(table, 0, 4, true)
	mainUi.StackPanel = stackPanel
}
