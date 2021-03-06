package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/mazegen"
	"golang.org/x/image/colornames"
)

type MazeDrawData struct {
	rows, columns                                    int
	originX, originY, cellSize, wallWidth, thickness float64
	drawWalls                                        bool
}

const windowHeight = 800
const windowWidth = 800

var mazeDrawData = &MazeDrawData{
	rows:      156,
	columns:   156,
	originX:   10,
	originY:   windowHeight - 10,
	cellSize:  5,
	wallWidth: 1,
	thickness: 0,
	drawWalls: true,
}

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

	debugPrintf := func(format string, a ...interface{}) {
		if debug {
			fmt.Printf(format, a...)
		}
	}

	debugPrintln := func(a ...interface{}) {
		if debug {
			fmt.Println(a)
		}
	}

	debugPrintf("debug=%t\n", debug)
	debugPrintf("drawWalls=%t\n", drawWalls)

	debugPrintln("initializing window...")
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
	debugPrintln("window initialized")

	state := "buildmaze"

	var grid *mazegen.Grid

	batch := pixel.NewBatch(&pixel.TrianglesData{}, nil)

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
			window.Clear(colornames.White)
			batch.Clear()
			drawMaze(batch, debug, window, grid, mazeDrawData)
			batch.Draw(window)
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
	batch *pixel.Batch,
	debug bool,
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
	{
		shape := imdraw.New(nil)
		shape.Color = colornames.Lightgray
		p0 := pixel.V(originX, originY)
		shape.Push(p0)

		width := float64(columns) * cellSize
		height := float64(rows) * cellSize
		shape.Push(pixel.V(originX+width, originY-height))

		shape.Rectangle(thickness)
		shape.Draw(window)
	}

	count := 0
	for y, cells := range grid.Cells {
		drawY := originY - (float64(y+1) * cellSize)
		for x, cell := range cells {

			drawX := originX + (float64(x) * cellSize)

			wallColor := colornames.Black

			if drawWalls {
				if cell.Walls[mazegen.North] {
					buildRectangle(drawX, drawY+(cellSize-wallWidth), wallWidth, cellSize, wallColor, thickness).Draw(batch)
				}
				if cell.Walls[mazegen.East] {
					buildRectangle(drawX+(cellSize-wallWidth), drawY, cellSize, wallWidth, wallColor, thickness).Draw(batch)
				}
				if cell.Walls[mazegen.South] {
					buildRectangle(drawX, drawY, wallWidth, cellSize, wallColor, thickness).Draw(batch)
				}
				if cell.Walls[mazegen.West] {
					buildRectangle(drawX, drawY, cellSize, wallWidth, wallColor, thickness).Draw(batch)
				}
			}

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

func main() {
	pixelgl.Run(run)
}
