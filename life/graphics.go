package life

import (
	"github.com/hajimehoshi/ebiten/v2"
    // "fmt"
    "time"
)

const (
    LOGICAL_FRACTION = 4
    LOGICAL_WIDTH_FRACTION = LOGICAL_FRACTION
    LOGICAL_HEIGHT_FRACTION = LOGICAL_FRACTION

    RED = 0
    GREEN = 1
    BLUE = 2
)

var (
    COLOR_NIGHT_GRAY = Color{R: 60, G: 60, B: 75}
)

type Game struct {
    life Life
    stepDelay time.Duration
    lastStepTime time.Time
    pixelsGrid [][][]byte
    pixels []byte
    logicalWidthFraction int
    logicalHeightFraction int
}

type Color struct {
    R, G, B byte
}

func (g *Game) Update() error {
    var now = time.Now()
    if (now.Sub(g.lastStepTime) >= g.stepDelay) {
        g.life.Next()
        g.lastStepTime = now
    }

    // fmt.Printf("FPS: %.2f\n", ebiten.ActualTPS())

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    for i := range g.life.grid {
        for j := range g.life.grid[i] {
            if g.life.grid[i][j] == true {
                var x = byte((255 / g.life.rows) * i)
                var y = byte((255 / g.life.cols) * j)
                writePixel(g.pixelsGrid[i][j], Color{R: 200, G: x, B: y})
            } else {
                writePixel(g.pixelsGrid[i][j], COLOR_NIGHT_GRAY)
            }
        }
    }

    screen.WritePixels(g.pixels)
}

func (g *Game) Layout(oudsiteWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return oudsiteWidth / g.logicalWidthFraction, outsideHeight / g.logicalHeightFraction
}

func (g *Game) logicalSize() (int, int) {
    var screenWidth, screenHeight = ebiten.WindowSize()
    return g.Layout(screenWidth, screenHeight)
}

func NewGame() Game {
    var game = Game {}

    game.lastStepTime = time.Now()
    game.stepDelay = 200 * time.Millisecond

    game.logicalWidthFraction = LOGICAL_FRACTION
    game.logicalHeightFraction = LOGICAL_FRACTION

    var cols, rows = game.logicalSize()
    game.life = NewLife(rows, cols)

    game.setPixles()

    return game
}

func (g *Game) setPixles() {
    var width, height = g.logicalSize()
    var pixels1d = make([]byte, 4 * width * height)
    var pixels2d = Fracture(pixels1d, height)
    var pixels3d = make([][][]byte, height)

    for i := range pixels2d {
        pixels3d[i] = Fracture(pixels2d[i], width)
    }

    g.pixels = pixels1d
    g.pixelsGrid = pixels3d
}

func writePixel(pixel []byte, color Color) {
    if len(pixel) != 4 {
        panic("writePixel should only take byte arrays of length 4")
    }

    pixel[RED] = color.R
    pixel[GREEN] = color.G
    pixel[BLUE] = color.B
}
