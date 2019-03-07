package main

import (
	"github.com/DoubleWB/game_demo/controller"
	"github.com/DoubleWB/game_demo/model"
	"github.com/DoubleWB/game_demo/view"
	"github.com/faiface/pixel/pixelgl"
)

func main() {

	model := model.RestartModel()
	view := view.View{}
	controller := controller.Controller{
		M: *model,
		V: view,
	}

	pixelgl.Run(controller.RunGame)
}
