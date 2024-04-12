package main

import (
	// "time"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jaceCallihoo/go-life/life"
	// "os"
)

const GRID_SIZE = 60
const GRID_ROWS = GRID_SIZE
const GRID_COLS = GRID_SIZE

const (
    gRID_ROWS = 50 
    gRID_COLS = 50
    SCALE = 8
)

func main() {

    // var life = Life.NewLife(GRID_ROWS, GRID_COLS)
    //
    // life.InsertGrid(Life.GRID4, 10, 10)
    //
    // life.Next()
    //
    // return

    /*
    game, err := life.NewGame(gRID_ROWS, gRID_COLS)
    if err != nil {
        log.Fatal(err)
    }
    */
    game := life.NewGame(gRID_ROWS, gRID_COLS)

    ebiten.SetWindowSize(gRID_COLS * SCALE, gRID_ROWS * SCALE)
    ebiten.SetWindowTitle("Jace: Game of Life")
    // Life.Demo8()
    if err := ebiten.RunGame(&game); err != nil {
	log.Fatal(err)
    }

    // var life = Life.NewLife(GRID_ROWS, GRID_COLS)
    //
    // var grid, err = Life.GridFromFile("./grids/7.txt")
    //
    // if err != nil {
    //     panic(err)
    // }
    //
    // life.InsertGrid(grid, 20, 10)
    //
    // var grid2, _ = Life.GridFromFile("./grids/4.txt")
    // life.InsertGrid(grid2, 25, 35)
    //
    // life.PrintGrid()
    //
    // for true {
    //     time.Sleep(200 * time.Millisecond)
    //     life.Next()
    //     life.PrintGrid()
    // }
}
