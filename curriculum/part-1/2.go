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
		screen.SetSize(80, 24)

		// The game loop.
		for {
			// Draw screen.
			screen.DrawString(0, 0, "Press any key to Quit.", goro.Style{})
			screen.Flush()

			// Handle events.
			switch screen.WaitEvent().(type) {
			case goro.EventKey:
				goro.Quit()
			case goro.EventQuit:
				return
			}
		}
	})
}
