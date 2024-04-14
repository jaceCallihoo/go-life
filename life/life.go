package life

import (
	"fmt"
	"os"
	"slices"
)

type Life struct {
    rows int
    cols int

    gridStates [][][]bool
    currentGridState int
}

func NewLife (rows int, cols int) Life {
    var life = Life {}

    life.rows = rows
    life.cols = cols

    life.currentGridState = 0
    life.gridStates = make([][][]bool, 2)
    for i := range life.gridStates {
        life.gridStates[i] = make([][]bool, rows)
        for j := range life.gridStates[i] {
            life.gridStates[i][j] = make([]bool, cols)
        }
    }

    return life
}

func (l Life) getCurrentGrid() [][]bool {
    return l.gridStates[l.currentGridState]
}


func (l *Life) SetNumGridStates(numGridStates int) {
    if numGridStates < 2 || numGridStates == len(l.gridStates) {
        return 
    } 


    split := l.currentGridState + 1
    leftSplit := l.gridStates[:split]
    rightSplit := l.gridStates[split:]


    if numGridStates < len(l.gridStates) {
        /*
        var diff = len(l.gridStates) - numGridStates

        var a = l.currentGridState + diff + 1
        var x = Min(len(l.gridStates), a)
        var rhs = l.gridStates[x:]

        var b = a - len(l.gridStates)
        var y = Max(0, b)
        var lhs = l.gridStates[y:l.currentGridState + 1]

        l.gridStates = append(rhs, lhs...)
        l.currentGridState = numGridStates - 1
        */

        diff := len(l.gridStates) - numGridStates
        rsReducedSize := Min(diff, len(rightSplit))
        lsReducedSize := Max(diff - len(rightSplit), 0)
        
        fmt.Println("rs", rsReducedSize)
        fmt.Printf("Min(%d (diff), %d (len(rightSplit) - 1))\n", diff, len(rightSplit) - 1)
        rightSplit = rightSplit[rsReducedSize:]

        // lsReducedSize := Max(diff - (len(l.gridStates) - l.currentGridState), 0)
        fmt.Println("ls", lsReducedSize)
        fmt.Printf("Max(%d (diff - len(l.gridStates)), 0)\n", diff - len(l.gridStates))
        leftSplit = leftSplit[lsReducedSize:]

        l.gridStates = slices.Concat(leftSplit, rightSplit)
        l.currentGridState = len(leftSplit) - 1
        /*
        diff := len(l.gridStates) - numGridStates

        lhsEnd := Min(len(l.gridStates), l.currentGridState + diff + 1)
        lhsNew := l.gridStates[l.currentGridState:lhsEnd]

        rhsEnd := Max(0, l.currentGridState - diff)
        rhsNew := l.gridStates[:rhsEnd]

        l.gridStates = append(lhsNew, rhsNew...)
        l.currentGridState = 0
        */

    } else {
        diff := numGridStates - len(l.gridStates)
        /*

        rhs := l.gridStates[l.currentGridState:]
        lhs := l.gridStates[:l.currentGridState]

        newGrids := make([][][]bool, diff)
        for i := range newGrids {
            newGrids[i] = make([][]bool, l.rows)
            for j := range newGrids[i] {
                newGrids[i][j] = make([]bool, l.cols)
            }
        }

        l.gridStates = append(lhs, rhs...)
        l.gridStates = append(l.gridStates, newGrids...) 
        */


        ////
        newGrids := make([][][]bool, diff)
        for i := range newGrids {
            newGrids[i] = make([][]bool, l.rows)
            for j := range newGrids[i] {
                newGrids[i][j] = make([]bool, l.cols)
            }
        }

        l.gridStates = slices.Concat(leftSplit, newGrids, rightSplit)
    }
}

func (l Life) InsertGrid(insert [][]bool, xOffset int, yOffset int) {
    grid := l.getCurrentGrid()
    for i := range insert {
        for j := range insert[i] {
            var iTarget = i + yOffset
            var jTarget = j + xOffset
            if iTarget >= 0 && iTarget < l.rows && jTarget >= 0 && jTarget < l.cols {
                grid[iTarget][jTarget] = insert[i][j]
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
    nextGridIndex := (l.currentGridState + 1) % len(l.gridStates)
    nextGrid := l.gridStates[nextGridIndex]

    for i := 0; i < l.rows; i++ {
        for j := 0; j < l.cols; j++ {
            nextGrid[i][j] = l.cellLivesNext(i, j)
        }
    }

    l.currentGridState = nextGridIndex
}

func (l Life) cellLivesNext(row int, col int) bool {
    live_neighbors := l.countLiveNeighbors(row, col)
    grid := l.getCurrentGrid()

    if grid[row][col] == true && (live_neighbors == 2 || live_neighbors == 3) {
        return true
    }

    if grid[row][col] == false && live_neighbors == 3 {
        return true
    }

    return false
}

func (l Life) countLiveNeighbors(row int, col int) int {
    var live_neighbors = 0
    grid := l.getCurrentGrid()

    // up
    if row - 1 >= 0 && grid[row - 1][col] == true {
        live_neighbors++
    }

    // down
    if row + 1 < l.rows && grid[row + 1][col] == true {
        live_neighbors++
    }

    // left
    if col - 1 >= 0 && grid[row][col - 1] == true {
        live_neighbors++
    }

    // right
    if col + 1 < l.cols && grid[row][col + 1] == true {
        live_neighbors++
    }

    // up + left
    if row - 1 >= 0 && col - 1 >= 0 && grid[row - 1][col - 1] == true {
        live_neighbors++
    }

    // up + right
    if row - 1 >= 0 && col + 1 < l.cols && grid[row - 1][col + 1] == true {
        live_neighbors++
    }

    // down + left
    if row + 1 < l.rows && col - 1 >= 0 && grid[row + 1][col - 1] == true {
        live_neighbors++
    }

    // down + right
    if row + 1 < l.rows && col + 1 < l.cols && grid[row + 1][col + 1] == true {
        live_neighbors++
    }

    return live_neighbors
}

func (l *Life) PrintGrid() {
    var separatorLen = l.cols + 1
    var rowLen = l.cols + 1
    var buffer = make([]byte, (l.rows * rowLen) + separatorLen)
    grid := l.getCurrentGrid()

    for i := 0; i < l.rows; i++ {
        for j := 0; j < l.cols; j++ {
            if grid[i][j]  == true {
                fmt.Println("found something")
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

