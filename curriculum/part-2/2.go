package main

import (
	"github.com/kettek/goro"
	"log"

	"myproject/entity"
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
		player := entity.NewEntity(screen.Columns/2, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorWhite})

		for {
			// Draw screen.
			screen.DrawRune(player.X, player.Y, player.Rune, player.Style)
			screen.Flush()
			screen.DrawRune(player.X, player.Y, ' ', goro.Style{})

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					player.Move(action.X, action.Y)
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}
		}
	})
}
