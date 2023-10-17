package life

import (
	"github.com/hajimehoshi/ebiten/v2"
    "time"
)

func Demo1() {
    var window_width = 1000
    var window_height = 1000
    ebiten.SetWindowSize(window_width, window_height)
    ebiten.SetWindowTitle("Conway's Game of Life")

    var game = Ptr(NewGame())

    game.stepDelay = 80 * time.Millisecond

    var gameCols, gameRows = game.logicalSize()
    game.life.InsertGrid(GRID12, gameCols / 2 - len(GRID12[0]) / 2, gameRows / 2 - len(GRID12) / 2)

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}

func Demo2() {
    var window_width = 1000
    var window_height = 900
    ebiten.SetWindowSize(window_width, window_height)
    ebiten.SetWindowTitle("Conway's Game of Life")

    var game = Ptr(NewGame())

    game.stepDelay = 80 * time.Millisecond

    var gameCols, gameRows = game.logicalSize()

    game.life.InsertGrid(GRID11, gameCols / 8, gameRows / 2)

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}

func Demo3() {
    var window_width = 600
    var window_height = 600
    ebiten.SetWindowSize(window_width, window_height)
    ebiten.SetWindowTitle("Conway's Game of Life")

    var game = Ptr(NewGame())

    var gameCols, gameRows = game.logicalSize()

    game.life.InsertGrid(GRID9, gameCols / 2 - len(GRID9[0]) / 2, gameRows / 2 - len(GRID9) / 2)

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}

func Demo4() {
    var window_width = 600
    var window_height = 400
    ebiten.SetWindowSize(window_width, window_height)
    ebiten.SetWindowTitle("Conway's Game of Life")

    var game = Ptr(NewGame())

    game.stepDelay = 200 * time.Millisecond

    var gameCols, gameRows = game.logicalSize()
    game.life.InsertGrid(GRID12, gameCols / 2 - len(GRID12[0]) / 2, gameRows / 2 - len(GRID12) / 2)

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}

func Demo5() {
    var window_width = 300
    var window_height = 300
    ebiten.SetWindowSize(window_width, window_height)
    ebiten.SetWindowTitle("Conway's Game of Life")

    var game = Ptr(NewGame())

    game.stepDelay = 200 * time.Millisecond

    var gameCols, gameRows = game.logicalSize()
    game.life.InsertGrid(GRID4, gameCols / 2 - len(GRID4[0]) / 2, gameRows / 2 - len(GRID4) / 2)

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
