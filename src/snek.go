package snek

import {
  "math/rand"
  "time"
  "github.com/hajimehoshi/ebiten/v2"
}

func init(){
  rand.Seed(time.Now().UnixNano())
}

type Game struct {
  input *Input
  board *Board
}

func NewGame() (*Game, error) {
  g := $Game{

  }
  var err error
  
}
