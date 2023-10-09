package life

import (
    "fmt"
    "strings"
    "os"
)

type Life struct {
    rows int
    cols int
    grid [][]bool
    grid_next [][]bool
}

func NewLife (rows int, cols int) Life {

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

func (l Life) Next() {
    for i := int(0); i < l.rows; i++ {
        for j := int(0); j < l.cols; j++ {
            l.updateNextCell(i, j)
        }
    }

    l.grid = l.grid_next
}

func (l Life) updateNextCell(row int, col int) {
    l.grid_next[row][col] = l.cellLivesNext(row, col)
}

func (l Life) cellLivesNext(row int, col int) bool {
    var live_neighbors = l.countLiveNeighbors(row, col)

    if l.grid[row][col] == true && (live_neighbors == 2 || live_neighbors == 3) {
        return true
    }

    if l.grid[row][col] == false && live_neighbors == 3 {
        return true
    }

    return false
}

func (l Life) countLiveNeighbors(row int, col int) int {
    var live_neighbors = 0

    // up
    if row - 1 >= 0 && l.grid[row - 1][col] == true {
        live_neighbors++
    }

    // down
    if row + 1 < l.rows && l.grid[row + 1][col] == true {
        live_neighbors++
    }

    // left
    if col - 1 >= 0 && l.grid[row][col - 1] == true {
        live_neighbors++
    }

    // right
    if col + 1 < l.cols && l.grid[row][col + 1] == true {
        live_neighbors++
    }

    // up + left
    if row - 1 >= 0 && col - 1 >= 0 && l.grid[row - 1][col - 1] == true {
        live_neighbors++
    }

    // up + right
    if row - 1 >= 0 && col + 1 < l.cols && l.grid[row - 1][col + 1] == true {
        live_neighbors++
    }

    // down + left
    if row + 1 < l.rows && col - 1 >= 0 && l.grid[row + 1][col - 1] == true {
        live_neighbors++
    }

    // down + right
    if row + 1 < l.rows && col + 1 < l.cols && l.grid[row + 1][col + 1] == true {
        live_neighbors++
    }

    return live_neighbors
}

func (l Life) PrintGrid() {
    var buffer = make([]byte, l.rows * (l.cols + 1))

    for i := int(0); i < l.rows; i++ {
        for j := int(0); j < l.cols; j++ {
            if l.grid[i][j]  == true {
                buffer[(i * (l.cols + 1)) + j] = byte('#')
            } else {
                buffer[(i * (l.cols + 1)) + j] = byte(' ')
            }
        }
        buffer[(i * (l.cols + 1)) + l.cols] = byte('\n')
    }

    os.Stdout.Write(buffer)
    fmt.Println(strings.Repeat("-", int(l.cols)))
}
