package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/opcodes"
)

func createOpcodePanelUI(operationPanel *tview.Flex, ops [][]string) *tview.Flex {
	table := tview.NewTable().SetBorders(false)

	for i, op := range ops {
		temp := opcodes.StringToOpcode[op[0]]
		fmt.Println(i, op, byte(temp.Op))
		table.SetCell(i, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))
		if len(op) > 1 {
			table.SetCell(i, 1, tview.NewTableCell(op[0]+" "+op[1]))
		} else {
			table.SetCell(i, 1, tview.NewTableCell(op[0]))
		}
		table.SetCell(i, 3, tview.NewTableCell(fmt.Sprintf("0x%x", temp.Op)))
	}
	table.SetSelectable(true, false).
		SetSelectedFunc(func(row int, column int) {

			selectedOps := ops[:row]
			fmt.Println(selectedOps)
			// Turn current cell (breakpoint) colorwhite
			table.GetCell(row, column).SetTextColor(tcell.ColorWhite)
		})

	operationPanel.AddItem(table, 0, 4, true)
	return operationPanel
}
