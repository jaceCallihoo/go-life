package life

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
    RED = 0
    GREEN = 1
    BLUE = 2

    SCALE = 4
)

var (
    COLOR_NIGHT_GRAY = Color{R: 60, G: 60, B: 75}
    COLOR_SPACE_BLACK = Color{R: 25, G: 25, B: 35}
    COLOR_CARBON = Color{R: 30, G: 30, B: 30}
    COLOR_MEDIUM_SKY = Color{R: 85, G: 120, B: 185}
    COLOR_KINDA_BLUE = Color{R: 55, G: 55, B: 95}

    CURRENT_GAME_IDX = 0
    GAME_PARAMS = []GameParam{
        {
            life: NewLife(50, 120),
            startingGrid: GRID1,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID2,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID3,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID4,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID5,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID6,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID7,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID8,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 50),
            startingGrid: GRID9,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID10,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 50),
            startingGrid: GRID11,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
        {
            life: NewLife(50, 120),
            startingGrid: GRID12,
            redChannelFunc: rowParabolic,
            greenChannelFunc: colParabolic,
            blueChannelFunc: flat200,
        },
    }
)

type GameParam struct {
    life Life
    redChannelFunc func(*Game, int, int) byte
    greenChannelFunc func(*Game, int, int) byte
    blueChannelFunc func(*Game, int, int) byte
    inactiveColor Color 
    startingGrid [][]bool
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
}

type Color struct {
    R, G, B byte
}

var x = 0
func (g *Game) Update() error {
    // optimization: add a sleep
    var now = time.Now()
    if now.Sub(g.lastStepTime) >= g.stepDelay {
        // fmt.Println("update")
        g.life.Next()
        g.lastStepTime = now
    }

    // todo: use a switch instead?
    // use up and down arrows to change speed
    // should the state be saved whene moving back to a "demo"?
    
    // if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
    x = x + 1
    fmt.Println(x)
    if x % 10 == 0 {
        fmt.Println("right")
        CURRENT_GAME_IDX++;
        next_game := GAME_PARAMS[CURRENT_GAME_IDX]
        g.life = next_game.life
        fmt.Println(len(g.pixels))
        g.setPixles()

        // todo: set scale in game params
        ebiten.SetWindowSize(g.life.cols * SCALE, g.life.rows * SCALE)

        time.Sleep(1 * time.Second)
        fmt.Println(len(g.pixels))
        // todo: also need to set the window size
        g.redChannelFunc = next_game.redChannelFunc
        g.greenChannelFunc = next_game.greenChannelFunc 
        g.blueChannelFunc = next_game.blueChannelFunc 
        g.inactiveColor = next_game.inactiveColor
    } else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {

    } else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
        fmt.Println("up")
        g.stepDelay -= 50 * time.Millisecond
        g.stepDelay = max(g.stepDelay, 0)
    } else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
        fmt.Println("down")
        g.stepDelay += 50 * time.Millisecond
    } else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
        // restart the game (set the grid state to the starting grid)
    } else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
        fmt.Println("escape")
        return ebiten.Termination
    }

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    width, height := ebiten.WindowSize()
    fmt.Println("height", height, "width", width, "cols", g.life.cols, "rows", g.life.rows)
    if height != g.life.rows * SCALE || width != g.life.cols * SCALE {
        fmt.Println("resizing in draw")
        ebiten.SetWindowSize(g.life.cols * SCALE, g.life.rows * SCALE)
    }
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

    fmt.Println("Writing pixels:", len(g.pixels))
    x, y := ebiten.WindowSize()
    fmt.Println("Window size:", x, y, x * y)
    screen_bounds := screen.Bounds()
    fmt.Println("screen bounds", screen_bounds.Max, screen_bounds.Min)
    screen.WritePixels(g.pixels)
}

func (g *Game) Layout(oudsiteWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return g.life.cols, g.life.rows
}

func NewGame(gameParamsIndex int) Game {
    var game = Game {}

    // initialize all different games
    for i := range(GAME_PARAMS) {
        fmt.Println(i)
    }

    gameParams := GAME_PARAMS[gameParamsIndex]
    rows := gameParams.life.rows
    cols := gameParams.life.cols
    game.life = NewLife(rows, cols)

    game.lastStepTime = time.Now()
    game.stepDelay = 200 * time.Millisecond

    game.redChannelFunc = rowParabolic
    game.greenChannelFunc = colParabolic
    game.blueChannelFunc = flat200
    game.inactiveColor = COLOR_SPACE_BLACK

    // ebiten.SetWindowSize(cols * scale, rows * scale)
    game.life.InsertGrid(GRID12, cols / 2 - len(GRID12[0]) / 2, rows / 2 - len(GRID12) / 2)

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

