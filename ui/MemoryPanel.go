package ui

import (
	"strings"

	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func (mainUi *MainUi) createMemoryPanel(evm *evm.Evm) { // Called on Init
	//.SetBorder(true).SetTitle("Memory")
	memoryPanel := tview.NewTextView()
	memoryPanel.SetBorder(true).SetTitle("Memory")
	mainUi.MemoryPanel = memoryPanel
}

func (mainUi *MainUi) updateMemoryPanel(evm *evm.Evm) {
	/*
		for i, data := range evm.Memory {
			mainUi.MemoryPanel.SetCell(i, 0, tview.NewTableCell(fmt.Sprintf(data)))
		}
	*/
	mainUi.MemoryPanel.SetText(strings.Join(evm.Memory, ""))
}
