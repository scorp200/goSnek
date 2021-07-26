package goSnek

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Dir() (Body, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyW) &&
			g.dir.y == 0 {
		return Body{x: 0, y: -1}, true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyS) &&
			g.dir.y == 0 {
		return Body{x: 0, y: 1}, true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyA) &&
			g.dir.x == 0 {
		return Body{x: -1, y: 0}, true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) &&
			g.dir.x == 0 {
		return Body{x: 1, y: 0}, true
	}

	return Body{}, false
}
