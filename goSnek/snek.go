package goSnek

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init(){
  rand.Seed(time.Now().UnixNano())
}

const (
  Width = 600
  Height = 600
  Scale = 60
)

type Game struct {
  body []Body
  pos Body
  food Body
  size int
  bodyImage *ebiten.Image
  op ebiten.DrawImageOptions
}

func NewGame() (*Game) {
  g := &Game{
    size: Width / Scale,
    pos: Body{
      x: Scale / 2,
      y: Scale / 2,
    },
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
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  g.op.GeoM.Reset()
  g.op.GeoM.Translate(float64(g.pos.x * g.size), float64(g.pos.y * g.size))
  screen.DrawImage(g.bodyImage, &g.op)
}
