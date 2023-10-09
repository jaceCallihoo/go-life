package main

import (
    "time"
    Life "github.com/jaceCallihoo/go-life/life"
)

const GRID_SIZE = 30
const GRID_ROWS = GRID_SIZE
const GRID_COLS = GRID_SIZE

func main() {

    var life = Life.NewLife(GRID_ROWS, GRID_COLS)
    life.PrintGrid()

    for true {
        time.Sleep(200 * time.Millisecond)
        life.Next()
        life.PrintGrid()
    }
}
