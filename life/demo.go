package life

var GAME_PARAMS = []GameParam{
    {
        life: NewLife(50, 50),
        startingGrid: GRID1,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_NIGHT_GRAY,
        scale: 6,
    },
    {
        life: NewLife(11, 11),
        startingGrid: GRID2,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_SPACE_BLACK,
        scale: 8,
    },
    {
        life: NewLife(10, 10),
        startingGrid: GRID3,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_CARBON,
        scale: 8,
    },
    {
        life: NewLife(19, 19),
        startingGrid: GRID4,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_MEDIUM_SKY,
        scale: 9,
    },
    {
        life: NewLife(50, 50),
        startingGrid: GRID5,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_KINDA_BLUE,
        scale: 6,
    },
    {
        life: NewLife(50, 50),
        startingGrid: GRID6,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_NIGHT_GRAY,
        scale: 6,
    },
    {
        life: NewLife(50, 50),
        startingGrid: GRID7,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_NIGHT_GRAY,
        scale: 7,
    },
    {
        life: NewLife(50, 50),
        startingGrid: GRID8,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_NIGHT_GRAY,
        scale: 7,
    },
    {
        life: NewLife(50, 50),
        histories: 7,
        startingGrid: GRID9,
        redChannelFunc: lifetimeRed,
        greenChannelFunc: lifetimeGreen,
        blueChannelFunc: lifetimeBlue,
        inactiveColor: COLOR_CARBON,
        scale: 7,
    },
    {
        life: NewLife(200, 240),
        startingGrid: GRID10,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_NIGHT_GRAY,
        scale: 4,
    },
    {
        life: NewLife(240, 200),
        startingGrid: GRID11,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_NIGHT_GRAY,
        scale: 4,
    },
    {
        life: NewLife(50, 50),
        startingGrid: GRID12,
        redChannelFunc: rowParabolic,
        greenChannelFunc: colParabolic,
        blueChannelFunc: flat200,
        inactiveColor: COLOR_NIGHT_GRAY,
        scale: 8,
    },
}


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
