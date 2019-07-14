package entity

// Flags represents multiple states stored as in a bitflag.
type Flags = int

// Our entity flags.
const (
	BlockMovement Flags = 1 << iota
)
