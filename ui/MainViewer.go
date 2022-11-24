package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
)

type MainUi struct {
	MemoryPanel         *tview.TextView
	StackPanel          *tview.Table
	OpcodePanel         tview.Primitive
	StackAndMemoryPanel tview.Primitive
	Modal               *tview.Modal
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

func createModalUI() *tview.Flex {

	table := tview.NewTable()
	textView := tview.NewTextView().SetDynamicColors(true).SetRegions(true)
	banner := `  
	_____   ____  __               
	| __\ \ / /  \/  |              
	| _| \ V /| |\/| |              
	|___| \_/ |_| _|_|              
	| __|_ ___ __| |___ _ _ ___ _ _ 
	| _|\ \ / '_ \ / _ \ '_/ -_) '_|
	|___/_\_\ .__/_\___/_| \___|_|  
    		|_|                     

	created by: rottaj
`

	fmt.Fprintf(textView, banner)

	table.SetCell(1, 0, tview.NewTableCell("space").SetTextColor(tcell.ColorGreen))
	table.SetCell(1, 1, tview.NewTableCell("Enter session")) /*.SetTextColor(tcell.ColorRed))*/
	table.SetCell(2, 0, tview.NewTableCell("c").SetTextColor(tcell.ColorRed))
	table.SetCell(2, 1, tview.NewTableCell("Clear session")) /*.SetTextColor(tcell.ColorRed))*/
	table.SetCell(3, 0, tview.NewTableCell("?").SetTextColor(tcell.ColorYellow))
	table.SetCell(3, 1, tview.NewTableCell("Toggle modal")) /*.SetTextColor(tcell.ColorRed))*/
	table.SetCell(4, 0, tview.NewTableCell("ctrl c").SetTextColor(tcell.ColorRed))
	table.SetCell(4, 1, tview.NewTableCell("Exit")) /*.SetTextColor(tcell.ColorRed))*/

	box := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(textView, 0, 40, true).
		AddItem(table, 0, 12, true)
	box.SetBorder(true)
	box.SetTitle("EVM Explorer")
	return box
}

func InitializeMainViewer(evm *evm.Evm) (app *tview.Application) {

	var mainUi MainUi
	/*
		box := createModalUI()

		modal := func(p tview.Primitive, width, height int) tview.Primitive {
			return tview.NewFlex().
				AddItem(nil, 0, 1, false).
				AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
					AddItem(nil, 0, 1, false).
					AddItem(p, height, 1, true).
					AddItem(nil, 0, 1, false), width, 1, true).
				AddItem(nil, 0, 1, false)
		}
	*/
	/*
		modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})
	*/

	app = tview.NewApplication()
	pages := tview.NewPages()
	mainUi.createOpcodePanel(evm)
	mainUi.createStackPanel(evm)
	mainUi.createMemoryPanel(evm)
	mainUi.createStackAndMemoryPanel()
	//`mainUi.createModal()
	mainUi.createMainLayout()
	pages.AddPage("main", mainUi.Layout, true, true)
	//pages.AddPage("modal", modal(box, 50, 20), true, true)
	app.SetRoot(pages, true).SetFocus(pages)
	return app // called in main func (init.go)
}
