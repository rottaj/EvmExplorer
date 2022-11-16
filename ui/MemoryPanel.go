package ui

import (
	"github.com/rivo/tview"
)

func createMemoryPanelUI(memoryPanel *tview.Flex, ops [][]string) *tview.Flex {

	text := tview.NewTextView().SetText("0") // Itialize 0000000000000000000000000000000000000000000000000000000000000000
	memoryPanel.AddItem(text, 0, 1, false)
	return memoryPanel
}
