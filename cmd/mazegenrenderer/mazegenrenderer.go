package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/miketmoore/mazegen"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
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

	fmt.Println("initializing text...")
	// Initialize text
	// orig := pixel.V(20, 50)
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(20, 50), atlas)
	// txt := text.New(orig, text.Atlas7x13)
	txt.Color = colornames.Green
	fmt.Println("text initialized")

	state := "buildmaze"

	var grid *mazegen.Grid

	for !win.Closed() {

		// Quit application when user input matches
		if win.JustPressed(pixelgl.KeyQ) {
			os.Exit(1)
		}

		switch state {
		case "buildmaze":
			rows := 3
			cols := 3
			random := mazegen.NewRandom()
			grid = nil
			grid, err = mazegen.BuildMaze(rows, cols, random)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			state = "render"
		case "render":
			win.Clear(colornames.Black)
			originX := 150.0
			originY := 400.0
			cellSize := 100.0
			wallWidth := 5.0
			thickness := 0.0

			//  y,x
			// mazegen data
			//   ______ ______ ______ ______
			//  |0,0   |0,1   |0,2   |0,3   |
			//  |      |      |      |      |
			//   ______ ______ ______ ______
			//  |1,0   |1,1   |1,2   |1,3   |
			//  |      |      |      |      |
			//   ______ ______ ______ ______
			//  |2,0   |2,1   |2,2   |2,3   |
			//  |      |      |      |      |
			//   ______ ______ ______ ______

			//  y,x
			// pixel coordinates
			//   ______ ______ ______ ______
			//  |2,0   |2,1   |2,2   |2,3   |
			//  |      |      |      |      |
			//   ______ ______ ______ ______
			//  |1,0   |1,1   |1,2   |1,3   |
			//  |      |      |      |      |
			//   ______ ______ ______ ______
			//  |0,0   |0,1   |0,2   |0,3   |
			//  |      |      |      |      |
			//   ______ ______ ______ ______

			for y, cells := range grid.Cells {
				drawY := originY - (float64(y) * cellSize)
				for x, cell := range cells {

					drawX := originX + (float64(x) * cellSize)
					rectShape := buildRectangle(drawX, drawY, cellSize, cellSize, colornames.White, thickness)
					rectShape.Draw(win)

					if cell.Walls[mazegen.North] {
						buildRectangle(drawX, drawY+(cellSize-wallWidth), wallWidth, cellSize, colornames.Blue, thickness).Draw(win)
					}
					if cell.Walls[mazegen.East] {
						buildRectangle(drawX+(cellSize-wallWidth), drawY, cellSize, wallWidth, colornames.Blue, thickness).Draw(win)
					}
					if cell.Walls[mazegen.South] {
						buildRectangle(drawX, drawY, wallWidth, cellSize, colornames.Blue, thickness).Draw(win)
					}
					if cell.Walls[mazegen.West] {
						buildRectangle(drawX, drawY, cellSize, wallWidth, colornames.Blue, thickness).Draw(win)
					}

					// northStr := "N"
					// eastStr := "E"
					// southStr := "S"
					// westStr := "W"

					// if !cell.Walls[mazegen.North] {
					// 	northStr = "_"
					// }
					// if !cell.Walls[mazegen.East] {
					// 	eastStr = "_"
					// }
					// if !cell.Walls[mazegen.South] {
					// 	southStr = "_"
					// }
					// if !cell.Walls[mazegen.West] {
					// 	westStr = "_"
					// }

					message := fmt.Sprintf("%d,%d", y, x)
					// message := fmt.Sprintf("%d,%d %s%s%s%s", y, x, northStr, eastStr, southStr, westStr)

					txt.Clear()
					// txt.Color = colornames.Green
					// fmt.Fprintln(txt, message)
					rect := pixel.R(drawX, drawY, drawX+cellSize, drawY+cellSize)
					// txt.Draw(win, pixel.IM.Moved(rect.Center().Sub(txt.Bounds().Center())))
					fmt.Fprintln(txt, message)

					cellCenter := rect.Center()
					vectorDiff := cellCenter.Sub(txt.Bounds().Center())

					// matrix := pixel.IM.Moved(rect.Center().Sub(txt.Bounds().Center()))
					// matrix := pixel.IM.Scaled(txt.Orig, 2)
					matrix := pixel.IM.Moved(vectorDiff)
					matrix = matrix.Scaled(txt.Orig, 1)
					txt.Draw(win, matrix)

				}
			}
			state = "view"
		case "view":
			if win.JustPressed(pixelgl.KeyEnter) {
				state = "buildmaze"
			}
		}

		win.Update()

	}
}

func buildRectangle(x, y, w, h float64, color color.RGBA, thickness float64) *imdraw.IMDraw {
	shape := imdraw.New(nil)
	shape.Color = color
	shape.Push(pixel.V(x, y))
	shape.Push(pixel.V(x+h, y+w))
	shape.Rectangle(thickness)
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
