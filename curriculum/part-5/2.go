package entity

import (
	"github.com/kettek/goro"
)

// Entity is a type for representing an active creature on a Tile.
type Entity struct {
	X, Y  int
	Rune  rune
	Style goro.Style
	Name  string
	Flags Flags
}

// Move moves the entity by a given amount.
func (b *Entity) Move(x, y int) {
	b.X += x
	b.Y += y
}

// NewEntity returns a pointer to a new populated Entity.
func NewEntity(x int, y int, r rune, style goro.Style, name string, flags Flags) *Entity {
	return &Entity{
		X:     x,
		Y:     y,
		Rune:  r,
		Style: style,
		Name:  name,
		Flags: flags,
	}
}

// FindEntityAtLocation finds and returns the first entity at x and y matching the provided flags. If none exists, it returns nil.
func FindEntityAtLocation(entities []*Entity, x, y int, checkMask Flags, matchFlags Flags) *Entity {
	for _, e := range entities {
		if (e.Flags&checkMask) == matchFlags && e.X == x && e.Y == y {
			return e
		}
	}
	return nil
}
