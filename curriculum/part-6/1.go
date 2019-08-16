package interfaces

import (
	"github.com/kettek/goro"
)

type Entity interface {
	Move(x, y int)
	DistanceTo(other Entity) float64
	// Basic property accessors
	X() int
	SetX(int)
	Y() int
	SetY(int)
	Rune() rune
	SetRune(rune)
	Style() goro.Style
	SetStyle(goro.Style)
	Name() string
	SetName(string)
	Flags() int
	SetFlags(int)
	// Actor and Fighter accessors
	Actor() Actor
	SetActor(Actor)
	Fighter() Fighter
	SetFighter(Fighter)
}
