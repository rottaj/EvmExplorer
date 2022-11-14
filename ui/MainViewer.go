package ui

import (
	"github.com/rivo/tview"
)

type OperationUI struct {
	app   *tview.Application
	panel *tview.Flex
}

func InitializeMainViewer( /*ops []string*/ ) (app *tview.Application) {

	app = tview.NewApplication()
	pages := tview.NewPages()

	operationUI := createOperationPanel(app)
	layout := createMainLayout(operationUI)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true)
	return app
}

func createOperationPanel(app *tview.Application) (operationPanel *tview.Flex) {
	operationPanel = tview.NewFlex().SetDirection(tview.FlexRow)

	return operationPanel
}

func createMainLayout(operationPanel tview.Primitive) (layout *tview.Flex) {
	///// Main Layout /////
	mainLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(operationPanel, 0, 4, false)

	info := tview.NewTextView()
	info.SetBorder(true)
	info.SetText("<Go EVM Explorer - a rottaj project>")
	info.SetTextAlign(tview.AlignCenter)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainLayout, 0, 20, true).
		AddItem(info, 3, 1, false)

	return layout
}
