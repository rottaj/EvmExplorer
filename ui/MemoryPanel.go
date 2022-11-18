package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func (mainUi *MainUi) createMemoryPanel(evm *evm.Evm) { // Called on Init
	memoryPanel := tview.NewTable().SetBorder(true).SetTitle("Memory")

	// Iterate through ops and build Stack

	// Check opcodes and build stack

	mainUi.MemoryPanel = memoryPanel
	//text := tview.NewTextView().SetText(fmt.Sprintf("%d", evm.Pc)) // Itialize
	//memoryPanel.AddItem(text, 0, 1, false)
}

func (mainUi *MainUi) updateMemoryPanel(evm *evm.Evm) {

}
