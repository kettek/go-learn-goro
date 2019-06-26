package main

import (
	"github.com/kettek/goro"
	"log"

	"myproject/entity"
	"myproject/mapping"
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
		mapWidth, mapHeight := 80, 24

		colors := map[string]goro.Color{
			"darkWall":   goro.Color{R: 0, G: 0, B: 100, A: 255},
			"darkGround": goro.Color{R: 50, G: 50, B: 150, A: 255},
		}

		gameMap := mapping.GameMap{
			Width:  mapWidth,
			Height: mapHeight,
		}

		gameMap.Initialize()

		player := entity.NewEntity(screen.Columns/2, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorWhite})
		npc := entity.NewEntity(screen.Columns/2-5, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorYellow})

		entities := []*entity.Entity{
			player,
			npc,
		}

		for {
			// Draw screen.
			DrawAll(screen, entities, gameMap, colors)
			ClearAll(screen, entities)

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					if !gameMap.IsBlocked(player.X+action.X, player.Y+action.Y) {
						player.Move(action.X, action.Y)
					}
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}
		}
	})
}
