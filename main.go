package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/scorp200/goSnek/goSnek"
)

func main(){
  game := goSnek.NewGame()
  ebiten.SetWindowSize(goSnek.Width, goSnek.Height)
  ebiten.SetWindowTitle("go Snek")
  ebiten.RunGame(game)
}

