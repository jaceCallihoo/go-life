package main

import (
    "time"
    "os"
    Life "github.com/jaceCallihoo/go-life/life"
    "fmt"
)

const GRID_SIZE = 30
const GRID_ROWS = GRID_SIZE
const GRID_COLS = GRID_SIZE

func main() {

    var life = Life.NewLife(GRID_ROWS, GRID_COLS)

    fmt.Println(rune('\n'))

    var grid, err = Life.GridFromFile("./grids/1.txt")

    if err != nil {
        os.Stderr.WriteString(err.Error() + "\n")
        os.Exit(1)
    }

    life.InsertGrid(grid, 20, 5)
    life.InsertGrid(grid, 0, 0)
    life.InsertGrid(grid, 0, 15)

    life.PrintGrid()

    for true {
        time.Sleep(200 * time.Millisecond)
        life.Next()
        life.PrintGrid()
    }
}
