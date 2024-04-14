package life

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
    COLOR_KINDA_BLUE = Color{R: 55, G: 55, B: 95}

    CURRENT_GAME_IDX = 0
)

type GameParam struct {
    life Life
    redChannelFunc func(*Game, int, int) byte
    greenChannelFunc func(*Game, int, int) byte
    blueChannelFunc func(*Game, int, int) byte
    inactiveColor Color 
    startingGrid [][]bool
    scale int
    histories int
}

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

    scale int

    skipDraw bool

    gameIndex int
}

type Color struct {
    R, G, B byte
}

func (g *Game) GetCols() int {
    return g.life.cols
}
func (g *Game) GetRows() int {
    return g.life.rows
}

func (g *Game) handleInput() {
    if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
        g.gameIndex++
        if g.gameIndex >= len(GAME_PARAMS) {
            g.gameIndex = 0
        }
        g.loadGame(g.gameIndex)
    } else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
        g.gameIndex--
        if g.gameIndex < 0 {
            g.gameIndex = len(GAME_PARAMS) - 1
        }
        g.loadGame(g.gameIndex)
    } else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
        g.stepDelay -= 50 * time.Millisecond
        g.stepDelay = max(g.stepDelay, 0)
    } else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
        g.stepDelay += 50 * time.Millisecond
    } else if inpututil.IsKeyJustPressed(ebiten.KeyR) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
        g.restartGame(g.gameIndex)
    } else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
        os.Exit(0)
    } else if inpututil.IsKeyJustPressed(ebiten.Key1) {
        fmt.Println("increase")
        g.life.SetNumGridStates(len(g.life.gridStates) + 1)
    } else if inpututil.IsKeyJustPressed(ebiten.Key2) {
        fmt.Println("reduce")
        g.life.SetNumGridStates(len(g.life.gridStates) - 1)
    }
}

func (g *Game) Update() error {
    now := time.Now()
    if now.Sub(g.lastStepTime) >= g.stepDelay {
        g.life.Next()
        g.lastStepTime = now
    }

    g.handleInput()

    return nil
}

func (g *Game) restartGame(gameParamsIndex int) {
    gp := &GAME_PARAMS[gameParamsIndex]
    gp.life = NewLife(g.life.rows, g.life.cols)
    gp.life.InsertGrid(gp.startingGrid, gp.life.cols / 2 - len(gp.startingGrid[0]) / 2, gp.life.rows / 2 - len(gp.startingGrid) / 2)
    g.life = gp.life
}

func (g *Game) Draw(screen *ebiten.Image) {
    if g.skipDraw {
        g.skipDraw = false
        return
    }

    for i := range g.life.grid {
        for j := range g.life.grid[i] {
            var color Color
            if g.life.grid[i][j] == true {
                color = Color {
                    R: g.redChannelFunc(g, i, j),
                    G: g.greenChannelFunc(g, i, j),
                    B: g.blueChannelFunc(g, i, j),
                }
            } else {
                color = g.inactiveColor
            }
            writePixel(g.pixelsGrid[i][j], color)
        }
    }

    screen.WritePixels(g.pixels)
}

func (g *Game) Layout(oudsiteWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return g.life.cols, g.life.rows
}

func NewGame(gameParamsIndex int) Game {
    game := Game {}

    for i := range(GAME_PARAMS) {
        gp := &GAME_PARAMS[i]
        gp.life.InsertGrid(gp.startingGrid, gp.life.cols / 2 - len(gp.startingGrid[0]) / 2, gp.life.rows / 2 - len(gp.startingGrid) / 2)
        fmt.Println(gp.inactiveColor)
    }

    game.lastStepTime = time.Now()
    game.stepDelay = 200 * time.Millisecond

    game.loadGame(0)

    return game
}

func (g *Game) loadGame(gameParamsIndex int) {
    gameParams := GAME_PARAMS[gameParamsIndex]

    if gameParams.life.cols != g.life.cols || gameParams.life.rows != g.life.rows {
        g.skipDraw = true
    }

    g.life = gameParams.life
    g.life.SetNumGridStates(gameParams.histories)
    g.redChannelFunc = gameParams.redChannelFunc
    g.greenChannelFunc = gameParams.greenChannelFunc
    g.blueChannelFunc = gameParams.blueChannelFunc 
    g.inactiveColor = gameParams.inactiveColor
    g.scale = gameParams.scale

    ebiten.SetWindowSize(g.life.cols * g.scale, g.life.rows * g.scale)
    g.setPixles()
}

func (g *Game) Run() {
    if err := ebiten.RunGame(g); err != nil {
        panic(err)
    }
}

func (g *Game) setPixles() {
    width, height := g.life.cols, g.life.rows
    pixels1d := make([]byte, 4 * width * height)
    pixels2d := Fracture(pixels1d, height)
    pixels3d := make([][][]byte, height)

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

