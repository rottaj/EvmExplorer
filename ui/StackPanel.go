package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func (mainUi *MainUi) createStackPanel(evm *evm.Evm) {
	stackPanel := tview.NewTable().SetBorder(true).SetTitle("Stack")
	mainUi.StackPanel = stackPanel
}
