package life

import (
    "testing"
)

func Test_InsertGrid(t *testing.T) {
    var tc = []struct {
        name string
        life Life
        grid [][]bool
        xOffset int
        yOffset int
        expect [][]bool
    } {
        {
            name: "Should insert a grid",
            life: NewLife(5, 5),
            grid: [][]bool {
                {true, true, true},
                {true, false, true},
                {true, true, false},
            },
            xOffset: 0,
            yOffset: 0,
            expect: [][]bool {
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
            grid: [][]bool {
                {true, true, true},
                {true, true, true},
                {true, true, true},
            },
            xOffset: 3,
            yOffset: 3,
            expect: [][]bool {
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
            grid: [][]bool {
                {true, true, true},
                {true, true, true},
                {true, true, true},
            },
            xOffset: -1,
            yOffset: -1,
            expect: [][]bool {
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
            grid: [][]bool {
                {true, true, true, false},
                {true, true, false, true},
                {true, false, true, true},
                {false, true, true, true},
            },
            xOffset: -1,
            yOffset: -1,
            expect: [][]bool {
                {true, false},
                {false, true},
            },
        },
    }

    for _, c := range tc {
        t.Run(c.name, func(t *testing.T) {
            c.life.InsertGrid(c.grid, c.xOffset, c.yOffset)
            if !equalSlice2d(c.life.grid, c.expect) {
                t.Error()
            }
        })
    }
}
