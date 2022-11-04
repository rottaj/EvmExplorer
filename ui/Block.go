package ui


import (
	"github.com/gdamore/tcell"
	//"github.com/ethereum/go-ethereum/core/types"
)



func (viewer *Viewer) Box(x1, y1, x2, y2 int, block string) {
  if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			viewer.Screen.SetContent(col, row, ' ', nil, viewer.BlockStyle)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		viewer.Screen.SetContent(col, y1, tcell.RuneHLine, nil, viewer.BlockStyle)
		viewer.Screen.SetContent(col, y2, tcell.RuneHLine, nil, viewer.BlockStyle)
	}
	for row := y1 + 1; row < y2; row++ {
		viewer.Screen.SetContent(x1, row, tcell.RuneVLine, nil, viewer.BlockStyle)
		viewer.Screen.SetContent(x2, row, tcell.RuneVLine, nil, viewer.BlockStyle)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		viewer.Screen.SetContent(x1, y1, tcell.RuneULCorner, nil, viewer.BlockStyle)
		viewer.Screen.SetContent(x2, y1, tcell.RuneURCorner, nil, viewer.BlockStyle)
		viewer.Screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, viewer.BlockStyle)
		viewer.Screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, viewer.BlockStyle)
	}

	//drawText(s, x1+1, y1+1, x2-1, y2-1, style, string(block.Time()))
	Text(viewer.Screen, x1+1, y1+1, x2-1, y2-1, viewer.DefaultStyle, block)
}
