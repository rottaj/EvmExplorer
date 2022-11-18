package ui

import (
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

type MainUi struct {
	MemoryPanel         *tview.Table
	StackPanel          *tview.Table
	OpcodePanel         tview.Primitive
	StackAndMemoryPanel tview.Primitive
	Layout              tview.Primitive
}

func (mainUi *MainUi) createStackAndMemoryPanel() {
	stackAndMemoryPanel := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(mainUi.StackPanel, 0, 20, false).
		AddItem(mainUi.MemoryPanel, 0, 20, false)
	mainUi.StackAndMemoryPanel = stackAndMemoryPanel
}

func (mainUi *MainUi) createMainLayout() {
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

func InitializeMainViewer(evm *evm.Evm) (app *tview.Application) {

	var mainUi MainUi

	app = tview.NewApplication()
	pages := tview.NewPages()

	mainUi.createOpcodePanel(evm)
	mainUi.createStackPanel(evm)
	mainUi.createMemoryPanel(evm)
	mainUi.createStackAndMemoryPanel()
	mainUi.createMainLayout()
	pages.AddPage("main", mainUi.Layout, true, true)

	app.SetRoot(pages, true).SetFocus(pages)
	return app // called in main func (init.go)
}
