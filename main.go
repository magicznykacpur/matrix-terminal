package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/term"
)

func main() {
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatalf("couldn't get terminal size: %v", err)
	}
	matrix := matrix{rows: height, cols: width, grid: newGrid(height, width), propagateTo: 1}

	ticker := time.NewTicker(time.Millisecond * 75)
	for {
		<-ticker.C
		fmt.Print("\033[H\033[2J")
		matrix.propagateRows()
		matrix.printMatrix()
	}
}
