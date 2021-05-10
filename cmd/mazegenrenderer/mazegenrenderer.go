package main

import (
	"fmt"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// func main() {
// 	rows := 2
// 	cols := 2
// 	random := mazegen.NewRandom()
// 	grid, err := mazegen.BuildMaze(rows, cols, random)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}

// 	for rowIndex, row := range grid.Cells {
// 		for columnIndex, cell := range row {
// 			fmt.Println(rowIndex, columnIndex, cell)
// 		}
// 	}
// }

func run() {

	// main, err := zelduh.NewMain(
	// 	debugMode,
	// 	tileSize,
	// 	frameRate,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	// err = main.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	// os.Exit(1)

	cellSize := 25.0

	// Initialize window
	fmt.Println("initializing window...")
	win, err := pixelgl.NewWindow(
		pixelgl.WindowConfig{
			Title:  "Maze",
			Bounds: pixel.R(0, 0, cellSize*100, cellSize*100),
			VSync:  true,
		},
	)
	if err != nil {
		fmt.Println("Initializing GUI window failed:")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("window initialized")

	for !win.Closed() {

		// Quit application when user input matches
		if win.JustPressed(pixelgl.KeyQ) {
			os.Exit(1)
		}

		win.Update()

	}
}

func main() {
	fmt.Println("main")
	pixelgl.Run(run)
}