package ui

import (
	"github.com/rivo/tview"
)

func MainViewer() {
	box := tview.NewBox().SetBorder(true).SetTitle("Go Eth Explorer")
	tview.NewApplication().SetRoot(box, true).Run()

}
