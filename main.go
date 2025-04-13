package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"
)

func main() {
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatalf("couldn't get terminal size: %v", err)
	}
	matrix := matrix{rows: height, cols: width, grid: newGrid(height, width)}

	fmt.Println("entering the matrix")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Millisecond * 75)
	for {
		select {
		case <-ticker.C:
			fmt.Print("\033[H\033[2J")
			matrix.propagateRows()
			matrix.printMatrix()
		case <-sigChan:
			fmt.Println("exitting the matrix.")
			os.Exit(0)
		}
	}
}
