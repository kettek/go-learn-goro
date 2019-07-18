package main

import (
	"myproject/interfaces"

	"github.com/kettek/goro/fov"
)

func InitializeFoV(gameMap interfaces.GameMap) fov.Map {
	fovMap := fov.NewMap(gameMap.Width(), gameMap.Height(), fov.AlgorithmBBQ)

	for x := 0; x < gameMap.Width(); x++ {
		for y := 0; y < gameMap.Height(); y++ {
			fovMap.SetBlocksMovement(x, y, gameMap.IsBlocked())
			fovMap.SetBlocksLight(x, y, gameMap.IsOpaque())
		}
	}

	return fovMap
}

func RecomputeFoV(fovMap fov.Map, centerX, centerY int, radius int, light fov.Light) {
	fovMap.Recompute(centerX, centerY, radius, light)
}
