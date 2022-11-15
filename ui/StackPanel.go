package ui

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
	"github.com/rottaj/EvmExplorer/evm"
	"github.com/rottaj/EvmExplorer/opcodes"
)

func createStackPanelUI(stackPanel *tview.Flex, ops [][]string) *tview.Flex {
	table := tview.NewTable().SetBorders(false)
	fmt.Println("TESTESTS", ops)

	// Iterate through ops and build Stack
	var stack evm.Stack

	// Check opcodes and build stack
	for i := 0; i <= len(ops)-1; i++ {
		opCode := opcodes.StringToOpcode[ops[i][0]]
		isPush := evm.IsPush(opCode)
		if isPush {

			fmt.Println("PUSH", i, ops[i], isPush)
			stack.Add(ops[i][1])
		} else {
			i++
		}

	}
	for i, x := range stack.StackValue {
		table.SetCell(i, 3, tview.NewTableCell(fmt.Sprintf(strconv.Itoa(i)+" "+x)))
	}

	stackPanel.AddItem(table, 0, 4, true)
	return stackPanel
}
