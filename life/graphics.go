package life

import (
	"github.com/hajimehoshi/ebiten/v2"
    "fmt"
    "time"
)

const (
    WINDOW_WIDTH = 400
    WINDOW_HEIGHT = 600

    LOGICAL_FRACTION = 4
    LOGICAL_WIDTH_FRACTION = LOGICAL_FRACTION
    LOGICAL_HEIGHT_FRACTION = LOGICAL_FRACTION

    RED = 0
    GREEN = 1
    BLUE = 2
)

type Game struct {
    life Life
    stepDelay time.Duration
    lastStepTime time.Time
    pixels [][][]byte
    pixelsFlat []byte
}

func (g *Game) Update() error {
    var now = time.Now()
    if (now.Sub(g.lastStepTime) >= g.stepDelay) {
        g.life.Next()
        g.lastStepTime = now
    }
    fmt.Println(ebiten.ActualTPS())

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    var width, height = g.LogicalSize()
    var pixels = make([]byte, 4 * width * height)
    var pixels2d = Fracture(pixels, height)
    var pixels3d = make([][][]byte, height)

    for i := range pixels2d {
        pixels3d[i] = Fracture(pixels2d[i], width)
    }

    for i := range g.life.grid {
        for j := range g.life.grid[i] {
            if g.life.grid[i][j] == true {
                pixels3d[i][j][RED] = 255
                pixels3d[i][j][BLUE] = 50
                pixels3d[i][j][GREEN] = 50
            } else {
                pixels3d[i][j][GREEN] = 100
                pixels3d[i][j][BLUE] = 100
            }
        }
    }

    // for i := range g.life.grid {
    //     for j := range g.life.grid[i] {
    //         if g.life.grid[i][j] == true {
    //             pixels[(i * g.life.cols + j) * 4] = 0xff
    //         } else {
    //             pixels[(i * g.life.cols + j) * 4 + 1] = 0xff
    //         }
    //     }
    // }

    screen.WritePixels(pixels)
}

func (g *Game) Layout(oudsiteWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return oudsiteWidth / LOGICAL_WIDTH_FRACTION, outsideHeight / LOGICAL_HEIGHT_FRACTION
}

func (g *Game) LogicalSize() (int, int) {
    var screenWidth, screenHeight = ebiten.WindowSize()
    return g.Layout(screenWidth, screenHeight)
}

func Test() {
    ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
    ebiten.SetWindowTitle("Conway's Game of Life")

    var game = ptr(NewGame())

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}


func NewGame() Game {

    // if screen size > || < max / min screen size, panic

    var game = Game {}

    var cols, rows = game.LogicalSize()
    game.life = NewLife(rows, cols)

    game.lastStepTime = time.Now()
    game.stepDelay = 200 * time.Millisecond

    var grid, _ = GridFromFile("./grids/7.txt")
    game.life.InsertGrid(grid, 20, 10)

    var grid2, _ = GridFromFile("./grids/4.txt")
    game.life.InsertGrid(grid2, 50, 20)

    // var grid3, _ = GridFromFile("./grids/8.txt")
    // game.life.InsertGrid(grid3, 50, 10)

    return game
}

func ptr[T any](val T) *T {
    return &val
}

func Fracture[T any](src []T, pieces int) [][]T {
    var fractured = make([][]T, pieces)
    var pieceLength = len(src) / pieces

    for i := range fractured {
        var pieceStart = i * pieceLength
        var pieceEnd = pieceStart + pieceLength
        fractured[i] = src[pieceStart:pieceEnd]
    }

    return fractured
}
