package goSnek

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	Width  = 600
	Height = 600
	Scale  = 60
	Fps    = 10
)

type Game struct {
	body       []Body
	pos        Body
	dir        Body
	food       Body
	size       int
	snakeSize  int
	bodyImage  *ebiten.Image
	op         ebiten.DrawImageOptions
	frameCount int
}

func NewGame() *Game {
	g := &Game{
		size: Width / Scale,
		pos: Body{
			x: Scale / 2,
			y: Scale / 2,
		},
		dir: Body{
			x: 1,
			y: 0,
		},
		snakeSize: 3,
	}
	g.bodyImage = ebiten.NewImage(g.size, g.size)
	g.bodyImage.Fill(color.White)
	return g
}

type Body struct {
	x int
	y int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func (g *Game) Update() error {
	g.frameCount++
	if g.frameCount >= 60/Fps {
		g.frameCount = 0

                g.pos = moveSnek(g.pos, g.dir)

		g.body = append(g.body, Body{x: g.pos.x, y: g.pos.y})

		for len(g.body) > g.snakeSize {
                  g.body = removeAt(g.body, 0)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.body {
		b := g.body[i]
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(b.x*g.size), float64(b.y*g.size))
		screen.DrawImage(g.bodyImage, &g.op)
	}
}
