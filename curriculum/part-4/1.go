package main

import (
	"myproject/mapping"

	"github.com/kettek/goro/fov"
)

func InitializeFoV(g *mapping.GameMap) fov.Map {
	fovMap := fov.NewMap(g.Width, g.Height, fov.AlgorithmBBQ)

	for x := range g.Tiles {
		for y, tile := range g.Tiles[x] {
			fovMap.SetBlocksMovement(x, y, tile.Flags&BlockMovement != 0)
			fovMap.SetBlocksLight(x, y, tile.Flags&tile.BlockSight != 0)
		}
	}

	return fovMap
}

func RecomputeFoV(fovMap fov.Map, centerX, centerY int, radius int, light fov.Light) {
	fovMap.Recompute(centerX, centerY, radius, light)
}
