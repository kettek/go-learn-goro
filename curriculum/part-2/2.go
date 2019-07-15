package entity

import (
	"myproject/interfaces"

	"github.com/kettek/goro"
)

// Entity is a type that represents an active entity in the world.
type Entity struct {
	x, y  int
	rune  rune
	style goro.Style
}

// NewEntity returns a pointer to a newly created Entity.
func NewEntity(x int, y int, r rune, s goro.Style) interfaces.Entity {
	return &Entity{
		x:     x,
		y:     y,
		rune:  r,
		style: s,
	}
}

// X returns the entity's X position.
func (e *Entity) X() int {
	return e.x
}

// SetX sets the entity's X position.
func (e *Entity) SetX(x int) {
	e.x = x
}

// Y returns the entity's Y position.
func (e *Entity) Y() int {
	return e.y
}

// SetY sets the entity's Y position.
func (e *Entity) SetX(y int) {
	e.y = y
}

// Rune returns the entity's rune.
func (e *Entity) Rune() rune {
	return e.rune
}

// SetRune sets the entity's rune.
func (e *Entity) SetRune(r rune) {
	e.rune = r
}

// Style returns the entity's goro.Style.
func (e *Entity) Style() goro.Style {
	return e.style
}

// SetStyle sets the entity's goro.Style.
func (e *Entity) SetStyle(s goro.Style) {
	e.style = s
}

// Move moves the entity by a given amount.
func (e *Entity) Move(x, y int) {
	e.X += x
	e.Y += y
}
