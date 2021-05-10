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

type MazeDrawData struct {
	rows, columns                                    int
	originX, originY, cellSize, wallWidth, thickness float64
	drawWalls                                        bool
}

const windowHeight = 800
const windowWidth = 800

func run() {

	argsWithoutProg := os.Args[1:]

	debug := false
	drawWalls := true
	if len(argsWithoutProg) > 0 {
		for _, value := range argsWithoutProg {
			if value == "debug" {
				debug = true
			}
		}
	}

	fmt.Printf("debug=%t\n", debug)
	fmt.Printf("drawWalls=%t\n", drawWalls)

	// Initialize window
	fmt.Println("initializing window...")
	window, err := pixelgl.NewWindow(
		pixelgl.WindowConfig{
			Title:  "Maze",
			Bounds: pixel.R(0, 0, windowWidth, windowHeight),
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
	text := text.New(pixel.V(20, 50), atlas)
	// text := text.New(orig, text.Atlas7x13)
	text.Color = colornames.Green
	fmt.Println("text initialized")

	state := "buildmaze"

	var grid *mazegen.Grid

	mazeDrawData := &MazeDrawData{
		rows:      15,
		columns:   15,
		originX:   10,
		originY:   windowHeight - 10,
		cellSize:  50,
		wallWidth: 2,
		thickness: 0,
		drawWalls: true,
	}

	for !window.Closed() {

		// Quit application when user input matches
		if window.JustPressed(pixelgl.KeyQ) {
			os.Exit(1)
		}

		switch state {
		case "buildmaze":
			random := mazegen.NewRandom()
			grid = nil
			grid, err = mazegen.BuildMaze(mazeDrawData.rows, mazeDrawData.columns, random)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			state = "render"
		case "render":
			window.Clear(colornames.Black)
			drawMaze(debug, text, window, grid, mazeDrawData)
			state = "view"
		case "view":
			if window.JustPressed(pixelgl.KeyEnter) {
				state = "buildmaze"
			}
		}

		window.Update()

	}
}

func drawMaze(
	debug bool,
	text *text.Text,
	window *pixelgl.Window,
	grid *mazegen.Grid,
	mazeDrawData *MazeDrawData,
) {

	rows := mazeDrawData.rows
	columns := mazeDrawData.columns
	originX := mazeDrawData.originX
	originY := mazeDrawData.originY
	cellSize := mazeDrawData.cellSize
	wallWidth := mazeDrawData.wallWidth
	thickness := mazeDrawData.thickness
	drawWalls := mazeDrawData.drawWalls

	// draw background
	// buildRectangle(
	// 	originX,
	// 	originY,
	// 	float64(columns)*cellSize,
	// 	cellSize,
	// 	// float64(rows)*cellSize,
	// 	// originX+(float64(columns)*cellSize),
	// 	// originY+(float64(rows)*cellSize),
	// 	colornames.White,
	// 	0,
	// ).Draw(window)
	{
		shape := imdraw.New(nil)
		shape.Color = colornames.White
		p0 := pixel.V(originX, originY)
		fmt.Println(p0)
		shape.Push(p0)

		width := float64(columns) * cellSize
		height := float64(rows) * cellSize
		shape.Push(pixel.V(originX+width, originY-height))

		shape.Rectangle(thickness)
		shape.Draw(window)
	}

	count := 0
	for y, cells := range grid.Cells {
		drawY := originY - (float64(y) * cellSize)
		for x, cell := range cells {

			drawX := originX + (float64(x) * cellSize)
			// rectShape := buildRectangle(drawX, drawY, cellSize, cellSize, colornames.White, 1)
			// rectShape.Draw(window)

			if drawWalls {
				if cell.Walls[mazegen.North] {
					buildRectangle(drawX, drawY+(cellSize-wallWidth), wallWidth, cellSize, colornames.Blue, thickness).Draw(window)
				}
				if cell.Walls[mazegen.East] {
					buildRectangle(drawX+(cellSize-wallWidth), drawY, cellSize, wallWidth, colornames.Blue, thickness).Draw(window)
				}
				if cell.Walls[mazegen.South] {
					buildRectangle(drawX, drawY, wallWidth, cellSize, colornames.Blue, thickness).Draw(window)
				}
				if cell.Walls[mazegen.West] {
					buildRectangle(drawX, drawY, cellSize, wallWidth, colornames.Blue, thickness).Draw(window)
				}
			}

			if debug {
				// message := fmt.Sprintf("%d,%d", y, x)
				message := fmt.Sprintf("%d", count)
				text.Clear()
				fmt.Fprintln(text, message)
			}

			// text.Color = colornames.Green
			// fmt.Fprintln(text, message)
			rect := pixel.R(drawX, drawY, drawX+cellSize, drawY+cellSize)
			// text.Draw(window, pixel.IM.Moved(rect.Center().Sub(text.Bounds().Center())))

			cellCenter := rect.Center()
			vectorDiff := cellCenter.Sub(text.Bounds().Center())

			// matrix := pixel.IM.Moved(rect.Center().Sub(text.Bounds().Center()))
			// matrix := pixel.IM.Scaled(text.Orig, 2)
			matrix := pixel.IM.Moved(vectorDiff)
			matrix = matrix.Scaled(text.Orig, 1)
			text.Draw(window, matrix)
			count++
		}
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
