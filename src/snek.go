package snek

import (
  "math/rand"
  "time"
)

func init(){
  rand.Seed(time.Now().UnixNano())
}

type Game struct {
}

func NewGame() (*Game, error) {
  g := &Game{

  }

  return g, nil
}
