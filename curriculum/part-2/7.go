package interfaces

// Our GameMap interface. This provides access to tile state and more.
type GameMap interface {
  Width() int
  Height() int
  IsBlocked(x, y int) bool
}
