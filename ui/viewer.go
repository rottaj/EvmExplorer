package ui

import (
  "log"
  "github.com/gdamore/tcell"
)

type Viewer struct {
}

func (viewer *Viewer) Init() {
  defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
  bannerStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
  //boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGrey)

  s, err := tcell.NewScreen()

  if err != nil {
    log.Fatalf("%+v", err)
  }
  if err := s.Init(); err != nil {
    log.Fatalf("%+v", err)
  }

  s.SetStyle(defStyle)
  s.EnableMouse()
  //s.EnablePaste()

  s.Clear()
  			//(x, y1, x2, y2)
  drawBanner(s, 1, 1, 80, 7, bannerStyle)
  //drawBox(s, 1, 10, 62, 40, boxStyle, "Press C to reset") // will pass in txn

  quit := func() {
    maybePanic := recover()
    s.Fini()
    if maybePanic != nil {
      panic(maybePanic)
    }
  }
  defer quit()
  

  //ox, oy := -1, -1
  for {

    // Update Screen
    s.Show()
    
    // Poll Event
    ev := s.PollEvent()

    // Process Event
    switch ev := ev.(type) {
      case *tcell.EventResize:
        s.Sync()
      case *tcell.EventKey:
        if ev.Key() == tcell.KeyCtrlC || ev.Key() == tcell.KeyEscape {
          return 
        } else if ev.Key() == tcell.KeyCtrlL {
          s.Sync()                
        } else if ev.Rune() == 'C' || ev.Rune() == 'c' {
          s.Clear() 
        }
    }
  }	
}


func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBanner(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
  if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	header := "Ethereum Block Explorer"
	headerTwo := "Created by: rottaj"


	drawText(s, x1+1, y1+1, x2-1, y2-1, style, header)
	drawText(s, x1+3, y1+3, x2-3, y2-3, style, headerTwo)
}

