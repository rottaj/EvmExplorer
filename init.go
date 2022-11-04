package main

import (
	//"fmt"
	_"log"
	"sync"
	_"time"
	"github.com/rottaj/GoEthExplorer/ui"
  	"github.com/rottaj/GoEthExplorer/util"
	_"github.com/rottaj/GoEthExplorer/eth"
	_"github.com/ethereum/go-ethereum/core/types"
	"github.com/gdamore/tcell"
)

func main() {

  	client := util.InitializeClient()
	_ = client

  	var viewer ui.Viewer
  	viewer.Init()




	quit := func() {
		maybePanic := recover()
		viewer.Screen.Fini()
		if maybePanic != nil {
		panic(maybePanic)
		}
	}
	defer quit()
	
	//ox, oy := -1, -1
	var wg sync.WaitGroup


	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Update Screen
			viewer.Screen.Show()
			// Poll Event
			ev := viewer.Screen.PollEvent()

			//Box(s, 1, 30, 62, 40, boxStyle, string(blocks.Time())) // will pass in txn

			switch ev := ev.(type) {
			case *tcell.EventResize:
				viewer.Screen.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyCtrlC || ev.Key() == tcell.KeyEscape {
				return 
				} else if ev.Key() == tcell.KeyCtrlL {
				viewer.Screen.Sync()                
				} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				viewer.Screen.Clear() 
				}
			}
		}	
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		y1, y2 := 10, 20
		for i := 0; i<= 10; i++ {
			//block := eth.GetLatestBlock()
			viewer.Box(1, y1, 62, y2, "Press C to reset") // will pass in txn
			y1 += 5
			y2 += 5
			//time.Sleep(12 * time.Second)
			
		}
	}()
	wg.Wait()
}



