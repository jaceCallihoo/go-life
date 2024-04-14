package life

import (
	"math"
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

func countCellLifetime(l Life, row, col int) int {
    var cellLifetime = 0
    for i := l.currentGridState; i >= 0; i-- {
        if l.gridStates[i][row][col] == false {
            return cellLifetime
        }
        cellLifetime++
    }

    for i := len(l.gridStates) - 1; i > l.currentGridState; i-- {
        if l.gridStates[i][row][col] == false {
            return cellLifetime
        }
        cellLifetime++
    }

    return cellLifetime
}

func lifetimeGreen(g *Game, row, col int) byte {
    var a = float64(len(g.life.gridStates) - 1)
    var cellLifetime = countCellLifetime(g.life, row, col)
    var x = float64(cellLifetime - 1)
    var y = x * (-1020.0 / a) + 1020
    return byte(math.Min(255.0, y))
}

func lifetimeRed(g *Game, row, col int) byte {
    var a = float64(len(g.life.gridStates) - 1)
    var cellLifetime = countCellLifetime(g.life, row, col)
    var x = float64(cellLifetime - 1)
    var y = x * (1020.0 / a)
    return byte(math.Min(255.0, y))
}

func lifetimeBlue(g *Game, row, col int) byte {
    var a = float64(len(g.life.gridStates) - 1)
    var cellLifetime = countCellLifetime(g.life, row, col)
    var x = float64(cellLifetime - 1)
    var y = x * (-60.0 / a) + 15
    return byte(math.Max(0.0, y))
}

