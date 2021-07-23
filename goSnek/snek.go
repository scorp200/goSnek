package goSnek

import (
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
)

type Game struct {
  body []Body
  pos Body
  food Body
  scale int
  size int
}

func NewGame() (*Game) {
  g := &Game{
  }

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

}
