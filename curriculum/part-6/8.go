package main

import (
	"fmt"
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
		maxMonstersPerRoom := 3
		gameState := PlayerTurnState

		fovRadius := 10

		colors := map[string]goro.Color{
			"darkWall":    goro.Color{R: 25, G: 25, B: 25, A: 255},
			"darkGround":  goro.Color{R: 100, G: 100, B: 100, A: 255},
			"lightWall":   goro.Color{R: 50, G: 50, B: 50, A: 255},
			"lightGround": goro.Color{R: 150, G: 150, B: 150, A: 255},
		}

		player := entity.NewPlayerEntity()

		entities := []interfaces.Entity{
			player,
		}

		gameMap := mapping.NewGameMap(mapWidth, mapHeight)

		gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, &entities, maxMonstersPerRoom)

		fovRecompute := true

		fovMap := InitializeFoV(gameMap)

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
					if gameState == PlayerTurnState {
						x := player.X() + action.X
						y := player.Y() + action.Y
						if !gameMap.IsBlocked(x, y) {
							otherEntity := entity.FindEntityAtLocation(entities, x, y, entity.BlockMovement, entity.BlockMovement)
							if otherEntity != nil {
								fmt.Printf("You lick the %s in the shins, much to its enjoyment!\n", otherEntity.Name())
							} else {
								player.Move(action.X, action.Y)
								fovRecompute = true
							}
						}
						gameState = NPCTurnState
					}
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}

			// Handle entity updates.
			if gameState == NPCTurnState {
				for _, e := range entities {
					if e.Actor() != nil {
						e.Actor().TakeTurn(fovMap, gameMap, entities)
					}
				}
				gameState = PlayerTurnState
			}

		}
	})
}
