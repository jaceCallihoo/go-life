package life

/*
import (
    "time"
    "math/rand"
)
func Demo1() {
    var game = NewGame(50, 120, 4)

    game.stepDelay = 80 * time.Millisecond

    var rows, cols = game.life.rows, game.life.cols
    game.life.InsertGrid(GRID12, cols / 2 - len(GRID12[0]) / 2, rows / 2 - len(GRID12) / 2)

    game.Run()
}

func Demo2() {
    var game = NewGame(50, 50, 4)

    game.stepDelay = 80 * time.Millisecond

    var rows, cols = game.life.rows, game.life.cols
    game.life.InsertGrid(GRID11, cols / 2, rows / 8)

    game.Run()
}

func Demo3() {
    var game = NewGame(50, 50, 4)

    var rows, cols = game.life.rows, game.life.cols
    game.life.InsertGrid(GRID9, cols / 2 - len(GRID9[0]) / 2, rows / 2 - len(GRID9) / 2)

    game.Run()
}

func Demo4() {
    var game = NewGame(50, 50, 4)

    game.stepDelay = 200 * time.Millisecond

    var rows, cols = game.life.rows, game.life.cols
    game.life.InsertGrid(GRID12, cols / 2 - len(GRID12[0]) / 2, rows / 2 - len(GRID12) / 2)

    game.Run()
}

func Demo5() {
    var game = NewGame(50, 50, 4)

    game.stepDelay = 200 * time.Millisecond

    var rows, cols = game.life.rows, game.life.cols
    game.life.InsertGrid(GRID4, cols / 2 - len(GRID4[0]) / 2, rows / 2 - len(GRID4) / 2)

    game.Run()
}

func Demo6() {
    var game = NewGame(50, 50, 4)

    game.redChannelFunc = rowSigmoid
    game.greenChannelFunc = colSigmoid
    game.blueChannelFunc = flat200

    game.stepDelay = 20 * time.Millisecond

    for i := range game.life.grid {
        for j := range game.life.grid[i] {
            if rand.Intn(2) == 1 {
                game.life.grid[i][j] = true
            }
        }
    }

    game.Run()
}

func Demo7() {
    var game = NewGame(50, 50, 4)

    game.stepDelay = 200000 * time.Millisecond

    game.redChannelFunc = rowSigmoid
    game.greenChannelFunc = colSigmoid
    game.blueChannelFunc = flat200

    for i := range game.life.grid {
        for j := range game.life.grid[i] {
            game.life.grid[i][j] = true
        }
    }

    game.Run()
}

func Demo8() {
    var game = NewGame(50, 50, 4)

    game.stepDelay = 100 * time.Millisecond

    game.redChannelFunc = lifetimeRed
    game.greenChannelFunc = lifetimeGreen
    game.blueChannelFunc = lifetimeBlue

    game.inactiveColor = COLOR_KINDA_BLUE

    for i := range game.life.grid {
        for j := range game.life.grid[i] {
            if rand.Intn(2) == 1 {
                game.life.grid[i][j] = true
            }
        }
    }

    game.Run()
}
*/
