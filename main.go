package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jaceCallihoo/go-life/life"
)

func main() {
	game := life.NewGame(0)

	ebiten.SetWindowTitle("Jace: Game of Life")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
