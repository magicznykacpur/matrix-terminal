package main

import (
	"fmt"
	"math/rand"
	"strings"

	color "github.com/fatih/color"
)

type matrix struct {
	rows        int
	cols        int
	grid        [][]rune
}

var letters = []rune(`ｦｱｳｴｵｶｷｹｺｻｼｽｾｿﾀﾂﾃﾅﾆﾇﾈﾊﾋﾎﾏﾐﾑﾒﾓﾔﾕﾗﾘﾜ:."=*+-¦|_`)

func newGrid(rows, cols int) [][]rune {
	grid := [][]rune{}

	for range rows {
		row := []rune{}
		for range cols {
			row = append(row, ' ')
		}
		grid = append(grid, row)
	}

	return grid
}

func (m *matrix) getRandomLetterRow() []rune {
	randomLetter := letters[rand.Intn(len(letters))]
	randIdx := rand.Intn(len(m.grid[0]))

	row := []rune{}
	for range m.cols {
		row = append(row, ' ')
	}

	row[randIdx] = rune(randomLetter)
	return row
}

func (m *matrix) propagateRows() {
	randLetterRow := m.getRandomLetterRow()
	newGrid := newGrid(m.rows, m.cols)

	newGrid[0] = randLetterRow

	for i := 1; i < len(m.grid); i++ {
		newGrid[i] = m.grid[i-1]
	}

	m.grid = newGrid
}

func (m *matrix) printMatrix() {
	out := ""

	for i := range m.grid {
		row := ""
		for _, r := range m.grid[i] {
			row += string(r)
		}
		if strings.TrimSpace(row) == "" {
			out += "\n"
		} else {
			out += row
		}
	}

	out = color.HiGreenString(out)
	fmt.Print(out)
}
