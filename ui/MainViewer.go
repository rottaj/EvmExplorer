package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

type MainUi struct {
	MemoryPanel         *tview.Box
	StackPanel          tview.Primitive
	OpcodePanel         tview.Primitive
	StackAndMemoryPanel tview.Primitive
	Layout              tview.Primitive
}

// Create Operations Panel - Opcodes Table

func (mainUi *MainUi) createStackAndMemoryPanel() {
	stackAndMemoryPanel := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(mainUi.StackPanel, 0, 20, false).
		AddItem(mainUi.MemoryPanel, 0, 20, false)
	mainUi.StackAndMemoryPanel = stackAndMemoryPanel
}

// Create Main Layout - Add Operations table to MainViewer
func (mainUi *MainUi) createMainLayout() {
	///// Main Layout /////

	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainUi.OpcodePanel, 0, 17, true).
		AddItem(mainUi.StackAndMemoryPanel, 0, 3, true)

	footer := tview.NewTextView()
	footer.SetBorder(true)
	footer.SetText("<EVM Explorer - a rottaj project>")
	footer.SetTextAlign(tview.AlignCenter)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainLayout, 0, 20, true).
		AddItem(footer, 3, 1, false)

	mainUi.Layout = layout
}

// MainViewer receives all opcodes, and empty stack.
// Passes opcodes & stack to child components.
func InitializeMainViewer(evm *evm.Evm) (app *tview.Application) {

	var mainUi MainUi

	app = tview.NewApplication()
	pages := tview.NewPages()

	mainUi.createOpcodePanel(evm)
	mainUi.createStackPanel(evm)
	mainUi.createMemoryPanel(evm)
	mainUi.createStackAndMemoryPanel()
	// call Everystate change
	//stackPanel := mainUi.createStackPanel(evm)   // Creates stackPanel (initalizes stack w/ pos 1)
	//memoryPanel := mainUi.createMemoryPanel(evm) // Creates MemoryPanel (initialized 32 bytes to 00)

	mainUi.createMainLayout()
	pages.AddPage("main", mainUi.Layout, true, true)

	app.SetRoot(pages, true).SetFocus(pages)
	return app // called in main func (init.go)
}
