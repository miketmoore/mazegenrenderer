package main

import (
	"fmt"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/mazegen"
	"golang.org/x/image/colornames"
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

	// Initialize window
	fmt.Println("initializing window...")
	win, err := pixelgl.NewWindow(
		pixelgl.WindowConfig{
			Title:  "Maze",
			Bounds: pixel.R(0, 0, 800, 800),
			VSync:  true,
		},
	)
	if err != nil {
		fmt.Println("Initializing GUI window failed:")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("window initialized")

	state := "buildmaze"

	var grid *mazegen.Grid

	for !win.Closed() {

		// Quit application when user input matches
		if win.JustPressed(pixelgl.KeyQ) {
			os.Exit(1)
		}

		switch state {
		case "buildmaze":
			rows := 20
			cols := 20
			random := mazegen.NewRandom()
			grid, err = mazegen.BuildMaze(rows, cols, random)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			state = "render"
		case "render":
			originX := 100.0
			originY := 100.0
			cellSize := 25.0
			for y, cells := range grid.Cells {
				drawY := originY + (float64(y) * cellSize)
				for x := range cells {
					shape := imdraw.New(nil)
					shape.Color = colornames.Blue

					drawX := originX + (float64(x) * cellSize)

					shape.Push(pixel.V(drawX, drawY))
					shape.Push(pixel.V(drawX+cellSize, drawY+cellSize))
					shape.Rectangle(1)
					shape.Draw(win)

				}
			}
			state = "view"
		case "view":
			//
		}

		win.Update()

	}
}

// func drawRectangle() {

// 	// rect := entity.componentShape.Shape
// 	// rect.Color = entity.componentColor.Color

// 	// rect.Push(entity.componentRectangle.Rect.Min)

// 	// rect.Push(entity.componentRectangle.Rect.Max)

// 	// rect.Rectangle(0)

// 	// rect.Draw(s.Win)
// }

func main() {
	fmt.Println("main")
	pixelgl.Run(run)
}
