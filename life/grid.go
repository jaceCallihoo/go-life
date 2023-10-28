package life

type Grid[T any] struct {
    data []T
    rows int
    cols int
}

func NewGrid[T any](rows, cols int) Grid[T] {
    var grid = Grid[T]{}
    grid.data = make([]T, rows * cols)
    grid.rows = rows
    grid.cols = cols
    return grid
}
