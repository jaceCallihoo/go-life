package life

import (
    "os"
    // "fmt"
)

type Life struct {
    rows int
    cols int
    grid [][]bool
    gridNext [][]bool

    gridStates [][][]bool
    currentGridState int
    numGridStates int
}

func NewLife (rows int, cols int) Life {
    var life = Life {}

    life.rows = rows
    life.cols = cols

    life.numGridStates = 2
    life.currentGridState = 0
    life.gridStates = make([][][]bool, life.numGridStates)
    for i := range life.gridStates {
        life.gridStates[i] = make([][]bool, rows)
        for j := range life.gridStates[i] {
            life.gridStates[i][j] = make([]bool, cols)
        }
    }

    life.grid = life.gridStates[0]
    life.gridNext = life.gridStates[1]


    return life
}

func (l Life) InsertGrid(grid [][]bool, xOffset int, yOffset int) {
    for i := range grid {
        for j := range grid[i] {
            var iTarget = i + yOffset
            var jTarget = j + xOffset
            if !(iTarget < 0 || iTarget >= l.rows || jTarget < 0 || jTarget >= l.cols) {
                l.grid[i + yOffset][j + xOffset] = grid[i][j]
            }
        }
    }
}

func GridFromFile(path string) ([][]bool, error)  {
    var data, err = os.ReadFile(path)

    if err != nil {
        return nil, err
    }

    var lines [][]byte
    var lineStart = 0
    for i := range data {
        if data[i] == '\n' {
            if i > 0 && data[i - 1] == '\r' {
                lines = append(lines, data[lineStart:i-1])
                lineStart = i + 1
            } else {
                lines = append(lines, data[lineStart:i])
                lineStart = i + 1
            }
        }
    }

    var grid = make([][]bool, len(lines))
    for i := range lines {
        var row = make([]bool, len(lines[i]))
        for j := range lines[i] {
            row[j] = lines[i][j] == '#'
        }
        grid[i] = row
    }

    return grid, nil
}

func (l *Life) Next() {
    for i := 0; i < l.rows; i++ {
        for j := 0; j < l.cols; j++ {
            l.updateNextCell(i, j)
        }
    }

    l.currentGridState = (l.currentGridState + 1) % l.numGridStates
    l.grid = l.gridStates[l.currentGridState]
    l.gridNext = l.gridStates[(l.currentGridState + 1) % l.numGridStates]
    // fmt.Println("     ", l.currentGridState)

    // l.grid, l.gridNext = l.gridNext, l.grid
}

func (l Life) updateNextCell(row int, col int) {
    l.gridNext[row][col] = l.cellLivesNext(row, col)
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
    var separatorLen = l.cols + 1
    var rowLen = l.cols + 1
    var buffer = make([]byte, (l.rows * rowLen) + separatorLen)

    for i := 0; i < l.rows; i++ {
        for j := 0; j < l.cols; j++ {
            if l.grid[i][j]  == true {
                buffer[(i * rowLen) + j] = '#'
            } else {
                buffer[(i * rowLen) + j] = ' '
            }
        }
        buffer[(i * (l.cols + 1)) + l.cols] = '\n'
    }

    for i := 0; i < separatorLen - 1; i++ {
        buffer[(l.rows * rowLen) + i] = '-'
    }

    buffer[len(buffer) - 1] = '\n'

    os.Stdout.Write(buffer)
}
