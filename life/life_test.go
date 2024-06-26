package life

import (
	"fmt"
	"testing"
)

func Test_InsertGrid(t *testing.T) {
	var tc = []struct {
		name    string
		life    Life
		grid    [][]bool
		xOffset int
		yOffset int
		expect  [][]bool
	}{
		{
			name: "Should insert a grid",
			life: NewLife(5, 5),
			grid: [][]bool{
				{true, true, true},
				{true, false, true},
				{true, true, false},
			},
			xOffset: 0,
			yOffset: 0,
			expect: [][]bool{
				{true, true, true, false, false},
				{true, false, true, false, false},
				{true, true, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
			},
		},
		{
			name: "Should allow grid to overlap bottom right boundry",
			life: NewLife(5, 5),
			grid: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			},
			xOffset: 3,
			yOffset: 3,
			expect: [][]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, true, true},
				{false, false, false, true, true},
			},
		},
		{
			name: "Should allow grid to overlap top left boundry",
			life: NewLife(5, 5),
			grid: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			},
			xOffset: -1,
			yOffset: -1,
			expect: [][]bool{
				{true, true, false, false, false},
				{true, true, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
			},
		},
		{
			name: "Should insert a grid that is bigger than the games grid",
			life: NewLife(2, 2),
			grid: [][]bool{
				{true, true, true, false},
				{true, true, false, true},
				{true, false, true, true},
				{false, true, true, true},
			},
			xOffset: -1,
			yOffset: -1,
			expect: [][]bool{
				{true, false},
				{false, true},
			},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			c.life.InsertGrid(c.grid, c.xOffset, c.yOffset)
			if !equalSlice2d(c.life.getCurrentGrid(), c.expect) {
				t.Error()
			}
		})
	}
}

func Test_cellLivesNext(t *testing.T) {
	var tc = []struct {
		name   string
		grid   [][]bool
		row    int
		col    int
		expect bool
	}{
		{
			name: "Should say that this live cell dies (of starvation)",
			grid: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			},
			row:    1,
			col:    1,
			expect: false,
		},
		{
			name: "Should say that this live cell dies (of overpopulation)",
			grid: [][]bool{
				{true, true, false},
				{false, true, true},
				{false, true, false},
			},
			row:    1,
			col:    1,
			expect: false,
		},
		{
			name: "Should say that this live cell lives (3 neighbors)",
			grid: [][]bool{
				{false, true, false},
				{false, true, false},
				{true, true, false},
			},
			row:    1,
			col:    1,
			expect: true,
		},
		{
			name: "Should say that this live cell lives (2 neighbors)",
			grid: [][]bool{
				{false, false, true},
				{false, true, false},
				{false, false, true},
			},
			row:    1,
			col:    1,
			expect: true,
		},
		{
			name: "Should say that this dead cell lives",
			grid: [][]bool{
				{false, false, true},
				{true, false, false},
				{true, false, false},
			},
			row:    1,
			col:    1,
			expect: true,
		},
		{
			name: "Should say that this dead corner cell lives",
			grid: [][]bool{
				{false, true, true},
				{true, true, false},
				{false, false, true},
			},
			row:    0,
			col:    0,
			expect: true,
		},
		{
			name: "Should say that this dead cell dies",
			grid: [][]bool{
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			row:    3,
			col:    2,
			expect: false,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			var life = NewLife(len(c.grid), len(c.grid[0]))
			life.InsertGrid(c.grid, 0, 0) 
			if life.cellLivesNext(c.row, c.col) != c.expect {
				t.Error()
			}
		})
	}
}

func Test_countLiveNeighbors(t *testing.T) {
	var tc = []struct {
		name   string
		grid   [][]bool
		row    int
		col    int
		expect int
	}{
		{
			name: "Should detect it has a neighbor",
			grid: [][]bool{
				{false, false, false},
				{true, false, false},
				{false, false, false},
			},
			row:    1,
			col:    1,
			expect: 1,
		},
		{
			name: "Should detect multiple neighbars",
			grid: [][]bool{
				{true, false, false},
				{false, false, true},
				{false, true, false},
			},
			row:    1,
			col:    1,
			expect: 3,
		},
		{
			name: "Should not count given cell towards the total",
			grid: [][]bool{
				{true, false, false},
				{false, true, false},
				{false, true, false},
			},
			row:    1,
			col:    1,
			expect: 2,
		},
		{
			name: "Should not count cells that are not neighbor",
			grid: [][]bool{
				{true, true, true, true, true},
				{true, false, false, false, true},
				{true, false, false, false, true},
				{true, false, false, false, true},
				{true, true, true, true, true},
			},
			row:    2,
			col:    2,
			expect: 0,
		},
		{
			name: "Should not access out of bounds cells",
			grid: [][]bool{
				{false},
			},
			row:    0,
			col:    0,
			expect: 0,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			var life = NewLife(len(c.grid), len(c.grid[0]))
			life.InsertGrid(c.grid, 0, 0)
			var val = life.countLiveNeighbors(c.row, c.col)
			if val != c.expect {
				t.Errorf("Expectend %d but recieved %d", c.expect, val)
			}
		})
	}
}

func Test_SetNumGridStates(t *testing.T) {
	var tc = []struct {
		name                   string
		gridStates             [][][]bool
		currentGridState       int
		expectGridStates       [][][]bool
	}{
		{
			name:             "Should remove the grid after the current grid state",
			currentGridState: 0,
			gridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
		},
		{
			name:             "Should remove the grid affter currentGridState when it's not 0",
			currentGridState: 1,
			gridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{false, true, false}, // 
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
			},
		},
		{
			name:             "Should remove multiple grids",
			currentGridState: 1,
			gridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{false, true, false}, //
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
			},
		},
		{
			name:             "Should remove grids before the current grid",
			currentGridState: 4,
			gridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false}, //
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{false, false, false}, //
					{false, true, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
			},
		},
		{
			name:             "Should remove grids both before and after the current grid",
			currentGridState: 3,
			gridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false}, //
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{false, false, false}, //
					{true, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
			},
		},
		{
			name:             "Should do nowthing when the input length is the same as current",
			currentGridState: 0,
			gridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
		},
		{
			name:             "Should insert empty grids",
			currentGridState: 0,
			gridStates: [][][]bool{
				{
					{true, false, false}, // 
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{true, false, false}, // 
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
		},
		{
			name:             "Should insert grid before the current",
			currentGridState: 1,
			gridStates: [][][]bool{
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, true, false}, //
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			expectGridStates: [][][]bool{
				{
					{false, true, false}, //
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
				{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
			},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			expectNumGridStates := len(c.expectGridStates)
			l := NewLife(len(c.gridStates[0]), len(c.gridStates[0][0]))
			l.currentGridState = c.currentGridState
			l.gridStates = c.gridStates

			l.SetNumGridStates(expectNumGridStates)

			if len(l.gridStates) != len(c.expectGridStates) {
				t.Fatalf("Unequal len(gridStates) - Recieved: %d, Expected: %d", len(l.gridStates), len(c.expectGridStates))
			}
			for i := range l.gridStates {
				for j := range l.gridStates[i] {
					for k := range l.gridStates[i][j] {
						if l.gridStates[i][j][k] {
							fmt.Print("#")
							// log.Printf("Live at (%d,%d) in index %d", j, k, i)
						} else {
							fmt.Print(".")
						}
					}
					fmt.Print("\n")
				}
				if i == l.currentGridState {
					fmt.Print(" ^\n")
				} else {
					fmt.Print("\n")
				}
			}

			fmt.Println("------------------------")
			for i := range c.expectGridStates {
				for j := range c.expectGridStates[i] {
					for k := range c.expectGridStates[i][j] {
						if c.expectGridStates[i][j][k] {
							fmt.Print("#")
							// log.Printf("Live at (%d,%d) in index %d", j, k, i)
						} else {
							fmt.Print(".")
						}
					}
					fmt.Print("\n")
				}
				if i == 0 {
					fmt.Print(" ^\n")
				} else {
					fmt.Print("\n")
				}
			}

			for i := range l.gridStates {
				nthOutputGridIndex := (i + l.currentGridState) % len(l.gridStates)
				nthExpectedGridIndex := i
				// fmt.Printf("Comapring output %d and expected %d", nthOutputGridIndex, nthExpectedGridIndex)
				if !equalSlice2d(l.gridStates[nthOutputGridIndex], c.expectGridStates[nthExpectedGridIndex]) {
					t.Errorf("Unequal gridState[%d]", i)
				}
			} 
		})
	}
}
