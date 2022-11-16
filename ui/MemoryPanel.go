package ui

import (
	"github.com/rivo/tview"
)

func createMemoryPanelUI(memoryPanel *tview.Flex, ops [][]string) *tview.Flex {

	text := tview.NewTextView().SetText("0000000000000000000000000000000000000000000000000000000000000040") // Itialize
	memoryPanel.AddItem(text, 0, 1, false)
	return memoryPanel
}
