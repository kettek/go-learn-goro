package mapping

import (
	"myproject/interfaces"
)

// GameMap is our map type for holding our tiles and dimensions.
type GameMap struct {
	width, height int
	tiles         [][]Tile
}

// NewGameMap initializes a GameMap's tiles to match the provided width and height and sets up a few tiles to block movement and sight. Returns a GameMap interface.
func NewGameMap(width, height int) interfaces.GameMap {
	g := &GameMap{
		width:  width,
		height: height,
	}
	g.tiles = make([][]Tile, g.width)

	for x := range g.tiles {
		g.tiles[x] = make([]Tile, g.height)
		for y := range g.tiles[x] {
			g.tiles[x][y].Flags = BlockMovement | BlockSight
		}
	}

	return g
}

// MakeMap populates our GameMap with rooms.
func (g *GameMap) MakeMap() {
	room1 := NewRect(20, 15, 10, 15)
	room2 := NewRect(35, 15, 10, 15)

	g.CreateRoom(room1)
	g.CreateRoom(room2)
}

// CreateRoom creates a room from a provided rect.
func (g *GameMap) CreateRoom(r Rect) {
	for x := r.X1 + 1; x < r.X2; x++ {
		for y := r.Y1 + 1; y < r.Y2; y++ {
			if g.InBounds(x, y) {
				g.tiles[x][y] = Tile{}
			}
		}
	}
}

// Width returns our GameMap's width.
func (g *GameMap) Width() int {
	return g.width
}

// Height returns our GameMap's height.
func (g *GameMap) Height() int {
	return g.height
}

// IsBlocked returns if the given coordinates are blocking movement.
func (g *GameMap) IsBlocked(x, y int) bool {
	// Always block if outside our GameMap's bounds.
	if !g.InBounds(x, y) {
		return true
	}
	return g.tiles[y][x].Flags&BlockMovement != 0
}

// InBounds returns if the given coordinates are within the map's bounds.
func (g *GameMap) InBounds(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}
	return true
}
