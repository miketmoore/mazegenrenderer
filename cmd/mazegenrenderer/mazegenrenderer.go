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

					drawX := originX + (float64(x) * cellSize)
					buildRectangle(drawX, drawY, cellSize, cellSize).Draw(win)

					// if cell.Walls[mazegen.North] == true {
					// 	// draw north wall

					// 	drawRectangle()
					// }

				}
			}
			state = "view"
		case "view":
			//
		}

		win.Update()

	}
}

func buildRectangle(x, y, w, h float64) *imdraw.IMDraw {
	shape := imdraw.New(nil)
	shape.Color = colornames.Blue
	shape.Push(pixel.V(x, y))
	shape.Push(pixel.V(x+h, y+w))
	shape.Rectangle(1)
	return shape
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
