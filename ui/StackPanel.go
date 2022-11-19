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
	stackPanel.SetCell(0, 0, tview.NewTableCell("POS"))
	stackPanel.SetCell(0, 1, tview.NewTableCell("DATA"))
	stackPanel.SetCell(0, 2, tview.NewTableCell("RAW"))
	for i, data := range evm.Stack {
		stackPanel.SetCell(i+1, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))
		//stackPanel.SetCell(i, 1, tview.NewTableCell(fmt.Sprintf("%x", data)))
		stackPanel.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprintf("%x", data)))
		stackPanel.SetCell(i+1, 2, tview.NewTableCell(data.String()))
	}
}
func (mainUi *MainUi) updateStackPanel(evm *evm.Evm) {
	//.SetBorder(true).SetTitle("Stack")
	mainUi.StackPanel.Clear()
	mainUi.StackPanel.SetCell(0, 0, tview.NewTableCell("POS"))
	mainUi.StackPanel.SetCell(0, 1, tview.NewTableCell("DATA"))
	mainUi.StackPanel.SetCell(0, 2, tview.NewTableCell("RAW"))
	for i, data := range evm.Stack {
		mainUi.StackPanel.SetCell(i+1, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))
		mainUi.StackPanel.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprintf("%x", data)))
		mainUi.StackPanel.SetCell(i+1, 2, tview.NewTableCell(data.String()))
	}
}
