package life

import (
	"github.com/hajimehoshi/ebiten/v2"
    "fmt"
)

const (
    CELL_WIDTH = 3
    CELL_HEIGHT = 3
    WINDOW_WIDTH = 200
    WINDOW_HEIGHT = 300
)

type Game struct {
    life Life
    images [][]*ebiten.Image
}

func (g *Game) Update() error {
    g.life.Next()
    fmt.Println(ebiten.ActualTPS())

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    var width, height = g.LogicalSize()
    var pixels = make([]byte, 4 * width * height)

    for i := range g.life.grid {
        for j := range g.life.grid[i] {
            if g.life.grid[i][j] == true {
                pixels[(i * g.life.cols + j) * 4] = 0xff
            }
        }
    }

    screen.WritePixels(pixels)
}

func (g *Game) Layout(oudsiteWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return oudsiteWidth / CELL_WIDTH, outsideHeight / CELL_HEIGHT
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
    var game = Game {}

    var cols, rows = game.LogicalSize()
    game.life = NewLife(rows, cols)

    game.images = make([][]*ebiten.Image, rows)
    for i := 0; i < rows; i++ {
       game.images[i] = make([]*ebiten.Image, cols)
       for j := 0; j < cols; j++ {
           var im = ebiten.NewImage(2, 2)
           // im.Fill(co)s
           game.images[i][j] = im
       }
    }

    var grid, _ = GridFromFile("./grids/7.txt")
    game.life.InsertGrid(grid, 20, 10)

    return game
}

func ptr[T any](val T) *T {
    return &val
}
