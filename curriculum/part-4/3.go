package main

import (
	"log"

	"myproject/entity"
	"myproject/interfaces"
	"myproject/mapping"

	"github.com/kettek/goro"
	"github.com/kettek/goro/fov"
)

func main() {
	// Initialize goro!
	if err := goro.InitTCell(); err != nil {
		log.Fatal(err)
	}

	goro.Run(func(screen *goro.Screen) {
		// Screen configuration.
		screen.SetTitle("goRo-game")

		// Randomize our seed so the map is randomized per run.
		goro.SetSeed(goro.RandomSeed())

		// Our initial variables.
		mapWidth, mapHeight := screen.Size()
		maxRooms, roomMinSize, roomMaxSize := 30, 6, 10

		fovRadius := 10
		fovRecompute := true

		colors := map[string]goro.Color{
			"darkWall":    goro.Color{R: 25, G: 25, B: 25, A: 255},
			"darkGround":  goro.Color{R: 100, G: 100, B: 100, A: 255},
			"lightWall":   goro.Color{R: 50, G: 50, B: 50, A: 255},
			"lightGround": goro.Color{R: 150, G: 150, B: 150, A: 255},
		}

		player := entity.NewEntity(screen.Columns/2, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorWhite})
		npc := entity.NewEntity(screen.Columns/2-5, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorYellow})

		entities := []interfaces.Entity{
			player,
			npc,
		}

		gameMap := mapping.NewGameMap(mapWidth, mapHeight)

		gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, player)

		fovMap := InitializeFoV(&gameMap)

		for {

			if fovRecompute {
				RecomputeFoV(fovMap, player.X(), player.Y(), fovRadius, fov.Light{})
			}

			// Draw screen.
			DrawAll(screen, entities, gameMap, fovMap, fovRecompute, colors)

			fovRecompute = false

			ClearAll(screen, entities, fovMap)

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					if !gameMap.IsBlocked(player.X()+action.X, player.Y()+action.Y) {
						player.Move(action.X, action.Y)
						fovRecompute = true
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
