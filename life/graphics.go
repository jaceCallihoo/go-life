package life

import (
    "github.com/hajimehoshi/ebiten/v2"
    "time"
)

const (
    RED = 0
    GREEN = 1
    BLUE = 2
)

var (
    COLOR_NIGHT_GRAY = Color{R: 60, G: 60, B: 75}
    COLOR_SPACE_BLACK = Color{R: 25, G: 25, B: 35}
    COLOR_CARBON = Color{R: 30, G: 30, B: 30}
    COLOR_MEDIUM_SKY = Color{R: 85, G: 120, B: 185}
    COLOR_KINDA_BLUE  = Color{R: 55, G: 55, B: 95}
)

type Game struct {
    life Life

    stepDelay time.Duration
    lastStepTime time.Time

    pixelsGrid [][][]byte
    pixels []byte

    redChannelFunc func(*Game, int, int) byte
    greenChannelFunc func(*Game, int, int) byte
    blueChannelFunc func(*Game, int, int) byte
    inactiveColor Color
}

type Color struct {
    R, G, B byte
}

func (g *Game) Update() error {
    var now = time.Now()
    if now.Sub(g.lastStepTime) >= g.stepDelay {
        g.life.Next()
        g.lastStepTime = now
    }

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    for i := range g.life.grid {
        for j := range g.life.grid[i] {
            if g.life.grid[i][j] == true {
                var color = Color {
                    R: g.redChannelFunc(g, i, j),
                    G: g.greenChannelFunc(g, i, j),
                    B: g.blueChannelFunc(g, i, j),
                }
                writePixel(g.pixelsGrid[i][j], color)
            } else {
                writePixel(g.pixelsGrid[i][j], g.inactiveColor)
            }
        }
    }

    screen.WritePixels(g.pixels)
}

func (g *Game) Layout(oudsiteWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return g.life.cols, g.life.rows
}

func NewGame(rows, cols, scale int) Game {
    var game = Game {}

    game.life = NewLife(rows, cols)

    game.lastStepTime = time.Now()
    game.stepDelay = 200 * time.Millisecond

    game.redChannelFunc = rowParabolic
    game.greenChannelFunc = colParabolic
    game.blueChannelFunc = flat200
    game.inactiveColor = COLOR_SPACE_BLACK

    ebiten.SetWindowSize(cols * scale, rows * scale)

    game.setPixles()

    return game
}

func (g *Game) Run() {
    if err := ebiten.RunGame(g); err != nil {
        panic(err)
    }
}

func (g *Game) setPixles() {
    var width, height = g.life.cols, g.life.rows
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

