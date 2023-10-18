package main

import (
    // "time"
    Life "github.com/jaceCallihoo/go-life/life"
    // "os"
)

const GRID_SIZE = 60
const GRID_ROWS = GRID_SIZE
const GRID_COLS = GRID_SIZE

func main() {

    // var life = Life.NewLife(GRID_ROWS, GRID_COLS)
    //
    // life.InsertGrid(Life.GRID4, 10, 10)
    //
    // life.Next()
    //
    // return

    Life.Demo2()

    return

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
