package main

import (
    "time"
    "os"
    Life "github.com/jaceCallihoo/go-life/life"
)

const GRID_SIZE = 30
const GRID_ROWS = GRID_SIZE
const GRID_COLS = GRID_SIZE

func main() {

    var life = Life.NewLife(GRID_ROWS, GRID_COLS)

    var grid, err = Life.GridFromFile("./grids/4.txt")

    if err != nil {
        os.Stderr.WriteString(err.Error() + "\n")
        os.Exit(1)
    }

    life.InsertGrid(grid, 10, 10)

    life.PrintGrid()

    for true {
        time.Sleep(200 * time.Millisecond)
        life.Next()
        life.PrintGrid()
    }
}
