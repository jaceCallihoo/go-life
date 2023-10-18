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
