package ui

import (
	"log"
	_"time"
	"github.com/gdamore/tcell"
)

type Viewer struct {
	Screen tcell.Screen
	DefaultStyle tcell.Style
	BannerStyle tcell.Style
	BlockStyle tcell.Style
}



func (viewer *Viewer) Init() {
	viewer.DefaultStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
  	viewer.BannerStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
  	viewer.BlockStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGrey)

	s, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	viewer.Screen = s

	viewer.Screen.SetStyle(viewer.DefaultStyle)
	viewer.Screen.EnableMouse()
	//s.EnablePaste()

	viewer.Screen.Clear()
		//(x, y1, x2, y2)
	Banner(viewer.Screen, 1, 1, 80, 7, viewer.BannerStyle)

}



