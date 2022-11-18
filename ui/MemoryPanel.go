package ui

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

func createMemoryPanelUI(memoryPanel *tview.Flex, evm *evm.Evm) *tview.Flex {

	//text := tview.NewTextView().SetText("0000000000000000000000000000000000000000000000000000000000000040") // Itialize
	text := tview.NewTextView().SetText(fmt.Sprintf("%d", evm.Pc)) // Itialize
	memoryPanel.AddItem(text, 0, 1, false)
	return memoryPanel
}

func updateMemoryPanelUI(memoryPanel *tview.Flex, evm *evm.Evm) {

	//text := tview.NewTextView().SetText("0000000000000000000000000000000000000000000000000000000000000040") // Itialize
	text := tview.NewTextView().SetText(fmt.Sprintf("%d", evm.Pc)) // Itialize
	memoryPanel.AddItem(text, 0, 1, false)
}
