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

func run() {

	debug := true
	drawWalls := false

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
			rows := 15
			cols := 15
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
			originX := 10.0
			originY := 740.0
			cellSize := 50.0
			wallWidth := 2.0
			thickness := 0.0

			for y, cells := range grid.Cells {
				drawY := originY - (float64(y) * cellSize)
				for x, cell := range cells {

					drawX := originX + (float64(x) * cellSize)
					rectShape := buildRectangle(drawX, drawY, cellSize, cellSize, colornames.White, 1)
					rectShape.Draw(win)

					if drawWalls {
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
					}

					if debug {
						message := fmt.Sprintf("%d,%d", y, x)
						txt.Clear()
						fmt.Fprintln(txt, message)
					}

					// txt.Color = colornames.Green
					// fmt.Fprintln(txt, message)
					rect := pixel.R(drawX, drawY, drawX+cellSize, drawY+cellSize)
					// txt.Draw(win, pixel.IM.Moved(rect.Center().Sub(txt.Bounds().Center())))

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
