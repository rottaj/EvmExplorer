package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
	//"github.com/rottaj/EvmExplorer/program"
)

func (mainUi *MainUi) createOpcodePanel(evm *evm.Evm) {
	opcodePanel := tview.NewFlex().SetDirection(tview.FlexRow)
	table := tview.NewTable().SetBorders(false).SetSeparator(tview.Borders.Vertical)

	table.SetCell(0, 0, tview.NewTableCell("STEP").SetTextColor(tcell.ColorRed).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("NAME").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 2, tview.NewTableCell("DATA").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 3, tview.NewTableCell("OPCODE").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 4, tview.NewTableCell("PC").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 5, tview.NewTableCell("GAS").SetTextColor(tcell.ColorYellow))
	for i, step := range evm.Steps {
		//currentGas += temp.StaticGas
		table.SetCell(i+1, 0, tview.NewTableCell("["+fmt.Sprint(i)+"]"))

		table.SetCell(i+1, 1, tview.NewTableCell(step.Mnemonic))
		if step.Data != nil {
			data := strings.Split(fmt.Sprintf("0x%x", step.Data), "")
			if len(data) > 9 { // Prevent rest of table from leaving screen
				data = data[:9]
				data = append(data, "...")

			}
			dataStr := strings.Join(data[:], "")
			table.SetCell(i+1, 2, tview.NewTableCell(dataStr))
			//table.SetCell(i+1, 2, tview.NewTableCell()
		} else {
			table.SetCell(i+1, 2, tview.NewTableCell(" "))
		}

		table.SetCell(i+1, 3, tview.NewTableCell(fmt.Sprintf("0x%x", step.Op)))
		table.SetCell(i+1, 4, tview.NewTableCell(fmt.Sprintf("%v", step.Pc)))
		table.SetCell(i+1, 5, tview.NewTableCell(fmt.Sprintf("%v", step.Gas)))
	}
	table.SetSelectable(true, false).
		SetSelectedFunc(func(row int, column int) {

			//selectedOps := evm.Ops[:row]
			//fmt.Println(selectedOps)
			// Turn current cell (breakpoint) colorwhite
			table.GetCell(row, column).SetTextColor(tcell.ColorWhite)
			evm.Debug(row)
			mainUi.StackPanel.Clear()
			mainUi.MemoryPanel.Clear()
			mainUi.updateStackPanel(evm)
			mainUi.updateMemoryPanel(evm)
			//mainUi.MemoryPanel.SetCell(i+1, 4, tview.NewTableCell(fmt.Sprintf("%v", currentGas)))
		})

	opcodePanel.AddItem(table, 0, 4, true)
	opcodePanel.SetBorder(true).SetTitle("Operations").SetTitleAlign(0)
	mainUi.OpcodePanel = opcodePanel
}
