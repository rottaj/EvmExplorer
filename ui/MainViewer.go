package ui

import (
	"github.com/rivo/tview"
)

// Create Operations Panel - Opcodes Table
func createOpcodePanel(opcodes [][]string) *tview.Flex {
	// opcodePanel is adds Table w/ opcodePanelAddItem in UI creator
	flexPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	opcodePanel := createOpcodePanelUI(flexPanel, opcodes)
	opcodePanel.SetBorder(true).SetTitle("Operations").SetTitleAlign(0)
	return opcodePanel
}

func createStackPanel(ops [][]string) *tview.Flex {
	flexPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	stackPanel := createStackPanelUI(flexPanel, ops)
	stackPanel.SetBorder(true).SetTitle("Stack").SetTitleAlign(0)
	return stackPanel
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
func InitializeMainViewer(ops [][]string) (app *tview.Application) {

	app = tview.NewApplication()
	pages := tview.NewPages()

	opcodePanel := createOpcodePanel(ops)   // Creates opcodePanel
	stackPanel := createStackPanel(ops[:5]) // Creates stackPanel (initalizes stack w/ pos 1)
	layout := createMainLayout(opcodePanel, stackPanel)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true).SetFocus(pages)
	return app
}
