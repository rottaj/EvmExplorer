package ui

import (
  "log"
  _"time"
  "os"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/gdamore/tcell"
  "github.com/joho/godotenv"
  _"github.com/rottaj/GoEthExplorer/eth"
)

type Viewer struct {
}

func (viewer *Viewer) Init() {

	err := godotenv.Load(".env") 
	if err != nil {
		log.Fatal(err)
	}
	RPC_URL := os.Getenv("RPC_URL")
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err) 
	}
	_ = client
	//fmt.Println("Client Connected")
	// Get Block
	/*
	block := eth.GetLatestBlock(client)
	_ = block
	time.Sleep(12 * time.Second)
	*/


	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
  	bannerStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
  	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGrey)

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
	drawBox(s, 1, 10, 62, 20, boxStyle, "Press C to reset") // will pass in txn

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

