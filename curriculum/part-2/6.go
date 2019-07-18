package mapping

// Tile represents the state of a given location in a GameMap.
type Tile struct {
	Flags uint
}

// Our Tile's flags.
const (
	BlockMovement = 1 << iota
	BlockSight
)
