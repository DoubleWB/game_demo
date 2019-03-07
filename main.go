package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"

	"github.com/DoubleWB/game_demo/controller"
	"github.com/DoubleWB/game_demo/model"
	"github.com/DoubleWB/game_demo/view"
)

func main() {
	//seed randomness
	rand.Seed(time.Now().UTC().UnixNano())
	//create model
	model := model.RestartModel()
	//create view
	view := view.View{}
	//unite with controler
	controller := controller.Controller{
		M: *model,
		V: view,
	}

	//run
	pixelgl.Run(controller.RunGame)
}
