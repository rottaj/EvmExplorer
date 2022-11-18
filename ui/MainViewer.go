package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

type MainUi struct {
	MemoryPanel tview.Primitive
	StackPanel  tview.Primitive
	OpcodePanel tview.Primitive
}

// Create Operations Panel - Opcodes Table
func createOpcodePanel(evm *evm.Evm) *tview.Flex {
	// opcodePanel is adds Table w/ opcodePanelAddItem in UI creator
	flexPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	opcodePanel := createOpcodePanelUI(flexPanel, evm)
	opcodePanel.SetBorder(true).SetTitle("Operations").SetTitleAlign(0)
	return opcodePanel
}

func createStackPanel(evm *evm.Evm) *tview.Flex { // Call every state change
	flexPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	stackPanel := createStackPanelUI(flexPanel, evm)
	stackPanel.SetBorder(true).SetTitle("Stack").SetTitleAlign(0)
	return stackPanel
}

func createMemoryPanel(evm *evm.Evm) *tview.Flex { // Call every state change
	flexPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	stackPanel := createMemoryPanelUI(flexPanel, evm)
	stackPanel.SetBorder(true).SetTitle("Memory").SetTitleAlign(0)
	return stackPanel
}

func createStackAndMemoryPanel(stackPanel tview.Primitive, memoryPanel tview.Primitive) *tview.Flex {
	stackAndMemoryPanel := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(stackPanel, 0, 20, false).
		AddItem(memoryPanel, 0, 20, false)
	return stackAndMemoryPanel
}

// Create Main Layout - Add Operations table to MainViewer
func createMainLayout(opcodePanel tview.Primitive, stackPanel tview.Primitive) (layout *tview.Flex) {
	///// Main Layout /////
	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(opcodePanel, 0, 17, true).
		AddItem(stackPanel, 0, 3, true)

	footer := tview.NewTextView()
	footer.SetBorder(true)
	footer.SetText("<EVM Explorer - a rottaj project>")
	footer.SetTextAlign(tview.AlignCenter)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainLayout, 0, 20, true).
		AddItem(footer, 3, 1, false)

	return layout
}

// MainViewer receives all opcodes, and empty stack.
// Passes opcodes & stack to child components.
func InitializeMainViewer(evm *evm.Evm) (app *tview.Application) {

	var mainUi MainUi

	app = tview.NewApplication()
	pages := tview.NewPages()

	opcodePanel := createOpcodePanel(evm) // Creates opcodePanel
	// call Everystate change
	stackPanel := createStackPanel(evm)   // Creates stackPanel (initalizes stack w/ pos 1)
	memoryPanel := createMemoryPanel(evm) // Creates MemoryPanel (initialized 32 bytes to 00)
	mainUi.MemoryPanel = memoryPanel
	mainUi.OpcodePanel = opcodePanel
	mainUi.StackPanel = stackPanel

	stackAndMemoryPanel := createStackAndMemoryPanel(stackPanel, memoryPanel)
	layout := createMainLayout(opcodePanel, stackAndMemoryPanel)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true).SetFocus(pages)
	return app // called in main func (init.go)
}
