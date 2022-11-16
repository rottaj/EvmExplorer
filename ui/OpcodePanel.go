package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/opcodes"
)

func createOpcodePanelUI(operationPanel *tview.Flex, ops [][]string) *tview.Flex {
	table := tview.NewTable().SetBorders(false).SetSeparator(tview.Borders.Vertical)

	var currentGas int = 21000 // Initialized to 21000
	table.SetCell(0, 0, tview.NewTableCell("STEP").SetTextColor(tcell.ColorYellow).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("NAME").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 2, tview.NewTableCell("PC").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 3, tview.NewTableCell("OPCODE").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 4, tview.NewTableCell("GAS").SetTextColor(tcell.ColorYellow))
	for i, op := range ops {
		temp := opcodes.StringToOpcode[op[0]]
		currentGas += temp.StaticGas
		fmt.Println(i+1, op, byte(temp.Op))
		table.SetCell(i+1, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))
		if len(op) > 1 {
			table.SetCell(i+1, 1, tview.NewTableCell(op[0]+" "+op[1]))
		} else {
			table.SetCell(i+1, 1, tview.NewTableCell(op[0]))
		}
		table.SetCell(i+1, 3, tview.NewTableCell(fmt.Sprintf("0x%x", temp.Op)))
		table.SetCell(i+1, 4, tview.NewTableCell(fmt.Sprintf("%v", currentGas)))
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
