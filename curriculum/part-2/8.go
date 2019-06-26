package main

import (
	"github.com/kettek/goro"

	"myproject/entity"
)

// DrawAll draws all entities to the screen and flushes it.
func DrawAll(screen *goro.Screen, entities []*entity.Entity) {
	for _, entity := range entities {
		DrawEntity(screen, entity)
	}
	screen.Flush()
}

// ClearAll clears all entities from the screen.
func ClearAll(screen *goro.Screen, entities []*entity.Entity) {
	for _, entity := range entities {
		ClearEntity(screen, entity)
	}
}

// DrawEntity draws a given entity to the screen.
func DrawEntity(screen *goro.Screen, e *entity.Entity) {
	screen.DrawRune(e.X, e.Y, e.Rune, e.Style)
}

// ClearEntity clears a given entity from the screen.
func ClearEntity(screen *goro.Screen, e *entity.Entity) {
	screen.DrawRune(e.X, e.Y, ' ', goro.Style{})
}
