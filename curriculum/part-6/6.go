package entity

import (
	"myproject/interfaces"

	"github.com/kettek/goro"
)

func NewPlayerEntity() interfaces.Entity {
	pc := &Entity{
		rune: '@',
		name: "Player",
		style: goro.Style{
			Foreground: goro.ColorWhite,
		},
		flags: BlockMovement,
	}
	return pc
}

func NewMonsterEntity(x, y int, r rune, style goro.Style, name string) interfaces.Entity {
	c := &Entity{
		x:     x,
		y:     y,
		rune:  r,
		name:  name,
		style: style,
		flags: BlockMovement,
	}
	c.SetActor(&MonsterActor{
		owner: c,
	})

	return c
}
