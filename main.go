package main

import (
	"log"
        "github.com/scorp200/GoSnek"
	"github.com/hajimehoshi/ebiten/v2"
)

func main(){
  game, err :=  snek.NewGame()
  if err != nil {
      log.Fatal(err)
  }
  ebiten.SetWindowSize(600,600)
  ebiten.SetWindowTitle("Go snek")
  if err := ebiten.RunGame(game); err != nil {
    log.Fatal(err)
  }
}
