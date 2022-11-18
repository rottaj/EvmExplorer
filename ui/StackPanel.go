package ui

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func (mainUi *MainUi) createStackPanel(evm *evm.Evm) {
	//.SetBorder(true).SetTitle("Stack")
	stackPanel := tview.NewTable()
	stackPanel.SetBorder(true).SetTitle("Stack")
	mainUi.StackPanel = stackPanel
	for i, data := range evm.Stack {
		stackPanel.SetCell(i, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))
		stackPanel.SetCell(i, 1, tview.NewTableCell(data.String()))
	}
}
func (mainUi *MainUi) updateStackPanel(evm *evm.Evm) {
	//.SetBorder(true).SetTitle("Stack")
	mainUi.StackPanel.Clear()
	for i, data := range evm.Stack {
		mainUi.StackPanel.SetCell(i, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))
		mainUi.StackPanel.SetCell(i, 1, tview.NewTableCell(data.String()))
	}
}
