package controller

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/DoubleWB/game_demo/model"
	"github.com/DoubleWB/game_demo/view"
)

type Controller struct {
	V view.View
	M model.Model
}

func (c Controller) RunGame() {
	cfg := pixelgl.WindowConfig{
		Title:  "Cutting Game Demo!",
		Bounds: pixel.R(0, 0, 1000, 1000),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		c.V.DrawToWindow(win, c.M)
		win.Update()
		if !c.M.Update(win) {
			return
		}
	}
}
