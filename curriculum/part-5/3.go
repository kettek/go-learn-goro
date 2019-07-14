package mapping

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro-game/entity"
)

// GameMap is our map type for holding our tiles and dimensions.
type GameMap struct {
	Width, Height int
	Tiles         [][]Tile
}

// Initialize initializes a GameMap's Tiles to match its Width and Height.
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

// MakeMap creates a new randomized map. This is built according to the passed arguments.
func (g *GameMap) MakeMap(maxRooms, roomMinSize, roomMaxSize int, entities *[]*entity.Entity, maxMonsters int) {
	var rooms []*Rect

	for r := 0; r < maxRooms; r++ {
		// Generate a random width and height.
		width := roomMinSize + goro.Random.Intn(roomMaxSize)
		height := roomMinSize + goro.Random.Intn(roomMaxSize)
		// Generate a random position within the map boundaries.
		x := goro.Random.Intn(g.Width - width - 1)
		y := goro.Random.Intn(g.Height - height - 1)
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
				(*entities)[0].X = roomCenterX
				(*entities)[0].Y = roomCenterY
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
			// Place random monsters in the room.
			g.PlaceEntities(room, entities, maxMonsters)
			// Append our new room to our rooms list.
			rooms = append(rooms, room)
		}
	}
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

// CreateHTunnel creates a horizontal tunnel from x1 to/from x1 starting at y.
func (g *GameMap) CreateHTunnel(x1, x2, y int) {
	for x := goro.MinInt(x1, x2); x <= goro.MaxInt(x1, x2); x++ {
		if g.InBounds(x, y) {
			g.Tiles[x][y] = Tile{}
		}
	}
}

// CreateVTunnel creates a vertical tunnel from y1 to/from y2 starting at x.
func (g *GameMap) CreateVTunnel(y1, y2, x int) {
	for y := goro.MinInt(y1, y2); y <= goro.MaxInt(y1, y2); y++ {
		if g.InBounds(x, y) {
			g.Tiles[x][y] = Tile{}
		}
	}
}

// PlaceEntities places 0 to maxMonsters monster entities in the provided room.
func (g *GameMap) PlaceEntities(room *Rect, entities *[]*entity.Entity, maxMonsters int) {
	monstersCount := goro.Random.Intn(maxMonsters)

	for i := 0; i < monstersCount; i++ {
		var monster *entity.Entity
		// Acquire a random location within the room.
		x := (1 + room.X1) + goro.Random.Intn(room.X2-room.X1-1)
		y := (1 + room.Y1) + goro.Random.Intn(room.Y2-room.Y1-1)

		if entity.FindEntityAtLocation(*entities, x, y, 0, 0) == nil {
			// Generate an orc with 80% probability or a troll with 20%.
			if goro.Random.Intn(100) < 80 {
				monster = entity.NewEntity(x, y, 'o', goro.Style{Foreground: goro.ColorLime}, "Orc", entity.BlockMovement)
			} else {
				monster = entity.NewEntity(x, y, 'T', goro.Style{Foreground: goro.ColorGreen}, "Troll", entity.BlockMovement)
			}
			*entities = append(*entities, monster)
		}
	}
}

// Explored returns if the tile at x by y has been explored.
func (g *GameMap) Explored(x, y int) bool {
	if g.InBounds(x, y) {
		return g.Tiles[x][y].Explored
	}
	return false
}

// SetExplored sets the explored state of the tile at x and y to the passed explored bool.
func (g *GameMap) SetExplored(x, y int, explored bool) {
	if g.InBounds(x, y) {
		g.Tiles[x][y].Explored = explored
	}
}

// IsBlocked returns if the given coordinates are blocking.
func (g *GameMap) IsBlocked(x, y int) bool {
	return g.Tiles[x][y].BlockMovement
}

// InBounds returns if the given coordinates are within the map's boundaries.
func (g *GameMap) InBounds(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return false
	}
	return true
}
