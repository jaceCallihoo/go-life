package life

import (
    "math"
    "fmt"
)

func rowLinear(g *Game, row, col int) byte {
    return byte(row * 255 / g.life.rows)
}

func colLinear(g *Game, row, col int) byte {
    return byte(col *  255 / g.life.cols)
}

func rowParabolic(g *Game, row, col int) byte {
    return byte(-((1020 * (row - g.life.rows) * row) / (g.life.rows * g.life.rows)))
}

func colParabolic(g *Game, row, col int) byte {
    return byte(-((1020 * (col - g.life.cols) * col) / (g.life.cols * g.life.cols)))
}

func flat200(g *Game, row, col int) byte {
    return byte(200)
}

func rowSigmoid(g *Game, row, col int) byte {
    var k = 0.06
    return byte(255 / (1 + math.Pow(math.E, -(k * float64(row - g.life.rows / 2)))))
}

func colSigmoid(g *Game, row, col int) byte {
    var k = 0.06
    return byte(255 / (1 + math.Pow(math.E, -(k * float64(col - g.life.cols / 2)))))
}

// func lifetimeRed(g *Game, row, col int) byte {
//     var cellLifetime = 0
//     for i := 0; i < g.life.numGridStates; i++ {
//         if g.life.gridStates[(g.life.currentGridState - i) % g.life.numGridStates][row][col] == false {
//             break
//         }
//         cellLifetime++
//     }
//
//     return 0
// }

func countCellLifetime(l Life, row, col int) int {
    var cellLifetime = 0
    for i := l.currentGridState; i >= 0; i-- {
        if l.gridStates[i][row][col] == false {
            return cellLifetime
        }
        cellLifetime++
    }

    for i := l.numGridStates - 1; i > l.currentGridState; i-- {
        if l.gridStates[i][row][col] == false {
            return cellLifetime
        }
        cellLifetime++
    }

    return cellLifetime
}

func lifetimeRed(g *Game, row, col int) byte {
    var cellLifetime = countCellLifetime(g.life, row, col)
    fmt.Println("cellLifetime: ", cellLifetime)
    return byte((cellLifetime - 1) * (255 / g.life.numGridStates))
}

func lifetimeGreen(g *Game, row, col int) byte {
    var cellLifetime = countCellLifetime(g.life, row, col)
    return byte((g.life.numGridStates - cellLifetime) * (255 / g.life.numGridStates))
}

func lifetimeBlue(g *Game, row, col int) byte {
    // var cellLifetime = countCellLifetime(g.life, row, col)
    return 15
}
