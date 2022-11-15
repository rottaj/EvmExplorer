package ui

import (
	"github.com/rivo/tview"
)

// Create Operations Panel - Opcodes Table
func createOperationStepsPanel(opcodes [][]string) (operationStepsPanel *tview.Flex) {
	flexPanel := tview.NewFlex().SetDirection(tview.FlexRow)

	opcodePanel := createOperationStepsUI(flexPanel, opcodes)
	opcodePanel.SetBorder(true).SetTitle("Operations").SetTitleAlign(0)

	return opcodePanel
}

// Create Main Layout - Add Operations table to MainViewer
func createMainLayout(opcodePanel tview.Primitive) (layout *tview.Flex) {
	///// Main Layout /////
	mainLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(opcodePanel, 0, 4, true)

	footer := tview.NewTextView()
	footer.SetBorder(true)
	footer.SetText("<EVM Explorer - a rottaj project>")
	footer.SetTextAlign(tview.AlignCenter)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainLayout, 0, 20, true).
		AddItem(footer, 3, 1, false)

	return layout
}

func InitializeMainViewer(ops [][]string) (app *tview.Application) {

	app = tview.NewApplication()
	pages := tview.NewPages()

	opcodePanel := createOperationStepsPanel(ops) // Creates Panel (calls OperationStepsUI)
	layout := createMainLayout(opcodePanel)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true).SetFocus(pages)
	return app
}
