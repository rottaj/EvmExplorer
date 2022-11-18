package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func (mainUi *MainUi) createMemoryPanel(evm *evm.Evm) { // Called on Init
	memoryPanel := tview.NewTable().SetBorder(true).SetTitle("Memory")
	mainUi.MemoryPanel = memoryPanel
}

func (mainUi *MainUi) updateMemoryPanel(evm *evm.Evm) {

}
