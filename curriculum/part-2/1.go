package interfaces

import (
	"github.com/kettek/goro"
)

type Entity interface {
	// Basic property accessors
	X() int
	SetX(int)
	Y() int
	SetY(int)
	Rune(rune)
	SetRune(rune)
	Style() goro.Style
	SetStyle(goro.Style)
	// More complex methods
}
