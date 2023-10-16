package life

import (
	"github.com/hajimehoshi/ebiten/v2"
    "image/color"
    "fmt"
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

    pixels[100] = 0xff
    pixels[101] = 0xff
    pixels[102] = 0xff

    for i := range g.life.grid {
        for j := range g.life.grid[i] {

        }
    }

    screen.WritePixels(pixels)

    return

    for i := range g.life.grid {
        for j := range g.life.grid[i] {
            var im = g.images[i][j]
            var co = color.RGBA{R: 255, G: uint8(i * 4), B: uint8(j * 4), A:255}
            if g.life.grid[i][j] == false {
                co = color.RGBA{}
            }
            im.Fill(co)
            var op = &ebiten.DrawImageOptions{}
            op.GeoM.Translate(float64(j) * 2, float64(i) * 2)
            screen.DrawImage(im, op)
        }
    }
}

func (g *Game) Layout(oudsiteWidth, outsideHeight int) (screenWidth, screenHeight int) {

    return oudsiteWidth, outsideHeight
}

func (g *Game) LogicalSize() (int, int) {
    var screenWidth, screenHeight = ebiten.WindowSize()
    return g.Layout(screenWidth, screenHeight)
}

func Test() {

    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Conway's Game of Life")

    var game = ptr(NewGame(60, 60))

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}


func NewGame(rows, cols int) Game {
    var game = Game {}

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
