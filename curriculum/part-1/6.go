package main

import (
	"github.com/kettek/goro"
	"log"
)

func main() {
	// Initialize goro!
	if err := goro.InitTCell(); err != nil {
		log.Fatal(err)
	}

	goro.Run(func(screen *goro.Screen) {
		// Screen configuration.
		screen.SetTitle("My Roguelike")

		// Our initial variables.
		playerX, playerY := screen.Columns/2, screen.Rows/2

		for {
			// Draw screen.
			screen.DrawRune(playerX, playerY, '@', goro.Style{})
			screen.Flush()
			screen.DrawRune(playerX, playerY, ' ', goro.Style{})

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					playerX += action.X
					playerY += action.Y
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}
		}
	})
}
