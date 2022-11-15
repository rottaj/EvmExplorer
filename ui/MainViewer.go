package ui

import (
	"github.com/rivo/tview"
)

type OperationUI struct {
	app   *tview.Application
	panel *tview.Flex
}

func InitializeMainViewer(ops [][]string) (app *tview.Application) {

	app = tview.NewApplication()
	pages := tview.NewPages()

	operationStepsPanel := createOperationStepsPanel(ops) // Creates Panel (calls OperationStepsUI)
	layout := createMainLayout(operationStepsPanel)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true)
	return app
}

func createOperationStepsPanel(opcodes [][]string) (operationStepsPanel *tview.Flex) {
	operationStepsPanel = tview.NewFlex().SetDirection(tview.FlexRow)

	operationsStepsPanel := createOperationStepsUI(operationStepsPanel, opcodes)
	operationsStepsPanel.SetBorder(true).SetTitle("Operations").SetTitleAlign(0)

	return operationStepsPanel
}

func createMainLayout(operationStepsPanel tview.Primitive) (layout *tview.Flex) {
	///// Main Layout /////
	mainLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(operationStepsPanel, 0, 4, false)

	footer := tview.NewTextView()
	footer.SetBorder(true)
	footer.SetText("<EVM Explorer - a rottaj project>")
	footer.SetTextAlign(tview.AlignCenter)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainLayout, 0, 20, true).
		AddItem(footer, 3, 1, false)

	return layout
}
