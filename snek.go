package goSnek

import (
	"math/rand"
	"time"
)

func init(){
  rand.Seed(time.Now().UnixNano())
}

type Game struct {
  body []Body
  pos Body
  food Body
  width int
  height int
  scale int
  size int 
}

func NewGame() (*Game, error) {
  g := &Game{

  }

  return g, nil
}

type Body struct {
  x int
  y int
}
