package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//Creates a board with w*h dimensions in the
//center of the window
func createBoard(size float64, window *pixelgl.Window) *imdraw.IMDraw {
	width, height := size, size
	center := window.Bounds().Center()
	board := imdraw.New(nil)

	board.Color = colornames.White
	board.Push(pixel.V(center.X - width / 2, center.Y + height / 2),
			   pixel.V(center.X + width / 2, center.Y + height / 2))
	board.Push(pixel.V(center.X - width / 2, center.Y - height / 2),
			   pixel.V(center.X + width / 2, center.Y - height / 2))
	board.Rectangle(0)

	return board
}

func createGrids(gridSize float64, boardSize float64, thick float64, window *pixelgl.Window) *imdraw.IMDraw {
	grid := imdraw.New(nil)
	
	grid.Color = colornames.Grey

	boardStartX := window.Bounds().Center().X - boardSize / 2
	boardStartY := window.Bounds().Center().Y - boardSize / 2

	boardEndX := window.Bounds().Center().X + boardSize / 2
	boardEndY := window.Bounds().Center().Y + boardSize / 2

	for x := int(boardStartX); x < int(boardEndX); x = x + int(gridSize - thick / 2) {
		grid.Push(pixel.V(float64(x) + thick, window.Bounds().Max.Y), pixel.V(float64(x) + thick, window.Bounds().Min.Y))
		grid.Line(thick)
	}
	
	for y := int(boardStartY); y < int(boardEndY); y = y + int(gridSize - thick / 2) {
		grid.Push(pixel.V(window.Bounds().Max.X, float64(y) + thick), pixel.V(window.Bounds().Min.X, float64(y) + thick))
		grid.Line(thick)
	}

	return grid
}

func run() {
	const WINDOW_HEIGHT = 640 
	const WINDOW_WIDTH = 960
	config := pixelgl.WindowConfig{
		Title: "Conway in Go",
		Bounds: pixel.R(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(config)

	if err != nil {
		panic(err)
	}

	board := createBoard(WINDOW_HEIGHT, win);
	grids := createGrids(29, WINDOW_HEIGHT, 1, win);

	for !win.Closed() {
		win.Clear(colornames.Grey)
		board.Draw(win)
		grids.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run);
}
