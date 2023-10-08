package main

import (
    "fmt"
    "strings"
    "time"
    "os"
)

const GRID_SIZE = 30
const GRID_ROWS = GRID_SIZE
const GRID_COLS = GRID_SIZE

var grid = [GRID_ROWS][GRID_COLS]bool{}
var next_grid = [GRID_ROWS][GRID_COLS]bool{}

type Life struct {
    rows uint32
    cols uint32
    grid [][]bool
    grid_next [][]bool
}

func newLife (rows uint32, cols uint32) Life {

    var life = Life{}

    life.rows = rows
    life.cols = cols
    
    var grid = make([][]bool, rows)
    for i := range grid {
       grid[i] = make([]bool, cols) 
    }
    life.grid = grid
    life.grid_next = grid
    
    return life
}

func main() {

    // var life = newLife(GRID_ROWS, GRID_COLS)

    grid[2][3] = true
    grid[2][4] = true
    grid[2][5] = true

    printGrid()

    for true {
        time.Sleep(200 * time.Millisecond)
        next()
        printGrid()
    }

}

func next() {
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            updateNextCell(i, j) 
        } 
    }

    grid = next_grid
}

func updateNextCell(row int, col int) {
    next_grid[row][col] = cellLivesNext(row, col)
}

func cellLivesNext(row int, col int) bool {
    var live_neighbors = countLiveNeighbors(row, col)

    if grid[row][col] == true && (live_neighbors == 2 || live_neighbors == 3) {
        return true 
    }

    if grid[row][col] == false && live_neighbors == 3 {
        return true
    }

    return false
}

func countLiveNeighbors(row int, col int) int {
    var live_neighbors = 0
 
    // up
    if row - 1 >= 0 && grid[row - 1][col] == true {
        live_neighbors++
    }
    
    // down
    if row + 1 < GRID_ROWS && grid[row + 1][col] == true {
        live_neighbors++
    }

    // left
    if col - 1 >= 0 && grid[row][col - 1] == true {
        live_neighbors++
    }

    // right
    if col + 1 < GRID_COLS && grid[row][col + 1] == true {
        live_neighbors++
    }

    // up + left
    if row - 1 >= 0 && col - 1 >= 0 && grid[row - 1][col - 1] == true {
        live_neighbors++
    }

    // up + right
    if row - 1 >= 0 && col + 1 < GRID_COLS && grid[row - 1][col + 1] == true {
        live_neighbors++
    }

    // down + left
    if row + 1 < GRID_ROWS && col - 1 >= 0 && grid[row + 1][col - 1] == true {
        live_neighbors++
    }

    // down + right
    if row + 1 < GRID_ROWS && col + 1 < GRID_COLS && grid[row + 1][col + 1] == true {
        live_neighbors++
    }

    return live_neighbors
}

func printGrid() {
    var buffer [GRID_ROWS * (GRID_COLS + 1)]byte

    for i := 0; i < GRID_ROWS; i++ {
        for j := 0; j < GRID_COLS; j++ {
            if grid[i][j]  == true {
                buffer[(i * (GRID_COLS + 1)) + j] = byte('#')
                // fmt.Print("█")
            } else {
                buffer[(i * (GRID_COLS + 1)) + j] = byte(' ')
                // fmt.Print(" ")
            }
        }
        // fmt.Print("\n")
        buffer[(i * (GRID_COLS + 1)) + GRID_COLS] = byte('\n')
    }
    // fmt.Print(buffer)
    os.Stdout.Write(buffer[:])
    fmt.Println(strings.Repeat("-", GRID_COLS))
}
