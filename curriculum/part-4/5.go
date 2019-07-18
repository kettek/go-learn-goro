package mapping

import (
	"mypackage/interfaces"

	"github.com/kettek/goro"
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

// MakeMap creates a new randomized map. This is built according to the passed arguments.
func (g *GameMap) MakeMap(maxRooms, roomMinSize, roomMaxSize int, player interfaces.Entity) {
	var rooms []Rect

	for r := 0; r < maxRooms; r++ {
		// Generate a random width and height.
		width := roomMinSize + goro.Random.Intn(roomMaxSize)
		height := roomMinSize + goro.Random.Intn(roomMaxSize)
		// Generate a random position within the map boundaries.
		x := goro.Random.Intn(g.width - width - 1)
		y := goro.Random.Intn(g.height - height - 1)
		// Create a Rect according to our generated sizes.
		room := NewRect(x, y, width, height)

		// Iterate through our existing rooms to check for intersection with our new room.
		intersects := false
		for _, otherRoom := range rooms {
			if room.Intersect(otherRoom) {
				intersects = true
				break
			}
		}
		// Add the room if there is no intersection found.
		if !intersects {
			g.CreateRoom(room)
			roomCenterX, roomCenterY := room.Center()

			// Always place the player in the center of the first room.
			if len(rooms) == 0 {
				player.SetX(roomCenterX)
				player.SetY(roomCenterY)
			} else {
				prevCenterX, prevCenterY := rooms[len(rooms)-1].Center()

				// Flip a coin if we should tunnel horizontally or vertically first.
				if goro.Random.Intn(1) == 1 {
					g.CreateHTunnel(prevCenterX, roomCenterX, prevCenterY)
					g.CreateVTunnel(prevCenterY, roomCenterY, roomCenterX)
				} else {
					g.CreateVTunnel(prevCenterY, roomCenterY, prevCenterX)
					g.CreateHTunnel(prevCenterX, roomCenterX, roomCenterY)
				}
			}
			// Append our new room to our rooms list.
			rooms = append(rooms, room)
		}
	}
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

// CreateHTunnel creates a horizontal tunnel from x1 to/from x1 starting at y.
func (g *GameMap) CreateHTunnel(x1, x2, y int) {
	for x := goro.MinInt(x1, x2); x <= goro.MaxInt(x1, x2); x++ {
		if g.InBounds(x, y) {
			g.tiles[x][y] = Tile{}
		}
	}
}

// CreateVTunnel creates a vertical tunnel from y1 to/from y2 starting at x.
func (g *GameMap) CreateVTunnel(y1, y2, x int) {
	for y := goro.MinInt(y1, y2); y <= goro.MaxInt(y1, y2); y++ {
		if g.InBounds(x, y) {
			g.tiles[x][y] = Tile{}
		}
	}
}

// Explored returns if the tile at x by y has been explored.
func (g *GameMap) Explored(x, y int) bool {
	if g.InBounds(x, y) {
		return g.tiles[x][y].Flags&Explored != 0
	}
	return false
}

// SetExplored sets the explored state of the tile at x and y to the passed explored bool.
func (g *GameMap) SetExplored(x, y int, explored bool) {
	if g.InBounds(x, y) {
		if explored {
			g.tiles[x][y].Flags = g.tiles[x][y].Flags | Explored
		} else {
			g.tiles[x][y].Flags = g.tiles[x][y].Flags &^ Explored
		}
	}
}

// IsBlocked returns if the given coordinates are blocking movement.
func (g *GameMap) IsBlocked(x, y int) bool {
	// Always block if outside our GameMap's bounds.
	if !g.InBounds(x, y) {
		return true
	}
	return g.tiles[x][y].Flags&BlockMovement != 0
}

// IsOpaque returns if the given coordinates are blocking sight.
func (g *GameMap) IsOpaque(x, y int) bool {
	if !g.InBounds(x, y) {
		return true
	}
	return g.tiles[x][y].Flags&BlockSight != 0
}

// InBounds returns if the given coordinates are within the map's boundaries.
func (g *GameMap) InBounds(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}
	return true
}
