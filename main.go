package main

import (
	"github.com/DoubleWB/game_demo/controller"
	"github.com/DoubleWB/game_demo/model"
	"github.com/DoubleWB/game_demo/view"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Cutting Game Demo!",
		Bounds: pixel.R(0, 0, 1000, 1000),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	blueX, blueY := 500.0, 700.0

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)

		if win.Pressed(pixelgl.KeyLeft) {
			blueX -= 10
		}

		if win.Pressed(pixelgl.KeyRight) {
			blueX += 10
		}

		if win.Pressed(pixelgl.KeyDown) {
			blueY -= 10
		}

		if win.Pressed(pixelgl.KeyUp) {
			blueY += 10
		}

		imd := imdraw.New(nil)

		imd.Color = pixel.RGB(1, 0, 0)
		imd.Push(pixel.V(200, 100))
		imd.Color = pixel.RGB(0, 1, 0)
		imd.Push(pixel.V(800, 100))
		imd.Color = pixel.RGB(0, 0, 1)
		imd.Push(pixel.V(blueX, blueY))
		imd.Polygon(0)
		imd.Draw(win)

		win.Update()
	}
}

func main() {

	model := model.RestartModel()
	view := view.View{}
	controller := controller.Controller{
		M: *model,
		V: view,
	}

	pixelgl.Run(controller.RunGame)
}
