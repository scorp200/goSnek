package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/scorp200/goSnek"
)

func main(){
  game = goSnek.NewGame()
  ebiten.SetWindowSize(game.width, game.height)
  ebiten.SetWindowTitle("go Snek") 
}

