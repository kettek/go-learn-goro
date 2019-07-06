package mapping

// GameMap is our map data type.
type GameMap struct {
	Width, Height int
	Tiles         [][]Tile
}

// Initialize initializes a GameMap's Tiles to match its Width and Height. It also sets up some coordinates to block movement and sight.
func (g *GameMap) Initialize() {
	g.Tiles = make([][]Tile, g.Width)

	for x := range g.Tiles {
		g.Tiles[x] = make([]Tile, g.Height)
		for y := range g.Tiles[x] {
			g.Tiles[x][y] = Tile{
				BlockSight:    true,
				BlockMovement: true,
			}
		}
	}
}

// MakeMap populates our GameMap with rooms.
func (g *GameMap) MakeMap() {
	room1 := NewRect(20, 15, 10, 15)
	room2 := NewRect(35, 15, 10, 15)

	g.CreateRoom(room1)
	g.CreateRoom(room2)
}

// CreateRoom creates a room from a provided rect.
func (g *GameMap) CreateRoom(r *Rect) {
	for x := r.X1 + 1; x < r.X2; x++ {
		for y := r.Y1 + 1; y < r.Y2; y++ {
			if g.InBounds(x, y) {
				g.Tiles[x][y] = Tile{}
			}
		}
	}
}

// IsBlocked returns if the given coordinates are blocking movement.
func (g *GameMap) IsBlocked(x, y int) bool {
	// Always block if outside our GameMap's bounds.
	if !g.InBounds(x, y) {
		return true
	}
	return g.Tiles[x][y].BlockMovement
}

// InBounds returns if the given coordinates are within the map's bounds.
func (g *GameMap) InBounds(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return false
	}
	return true
}
