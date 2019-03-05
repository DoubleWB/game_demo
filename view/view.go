package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/DoubleWB/game_demo/model"
	"github.com/DoubleWB/game_demo/util"
)

type View struct{}

func (v View) DrawToWindow(win *pixelgl.Window, m model.Model) {
	//draw bounding box
	bbox := imdraw.New(nil)
	bbox.Color = pixel.RGB(0, 0, 0)
	bbox.Push(pixel.Vec{
		X: util.BBOX_CORNERX,
		Y: util.BBOX_CORNERY,
	})
	bbox.Push(pixel.Vec{
		X: util.BBOX_DIM + util.BBOX_CORNERX,
		Y: util.BBOX_DIM + util.BBOX_CORNERY,
	})
	bbox.Rectangle(2.0)

	//draw enemy hp
	eHP := imdraw.New(nil)
	eHP.Color = pixel.RGB(1, 0, 0)
	eHP.Push(pixel.Vec{
		X: util.EHP_CORNERX,
		Y: util.EHP_CORNERY,
	})
	eHP.Push(pixel.Vec{
		X: util.BAR_WIDTH + util.EHP_CORNERX,
		Y: util.BAR_HEIGHT + util.EHP_CORNERY,
	})
	eHP.Rectangle(0)
	eHP.Color = pixel.RGB(0, 1, 0)
	eHP.Push(pixel.Vec{
		X: util.EHP_CORNERX,
		Y: util.EHP_CORNERY,
	})
	eHP.Push(pixel.Vec{
		X: (util.BAR_WIDTH * (m.EnemyHP / util.MAX_HP)) + util.EHP_CORNERX,
		Y: util.BAR_HEIGHT + util.EHP_CORNERY,
	})
	eHP.Rectangle(0)

	//draw player hp
	HP := imdraw.New(nil)
	HP.Color = pixel.RGB(1, 0, 0)
	HP.Push(pixel.Vec{
		X: util.HP_CORNERX,
		Y: util.HP_CORNERY,
	})
	HP.Push(pixel.Vec{
		X: util.BAR_WIDTH + util.HP_CORNERX,
		Y: util.BAR_HEIGHT + util.HP_CORNERY,
	})
	HP.Rectangle(0)
	HP.Color = pixel.RGB(0, 1, 0)
	HP.Push(pixel.Vec{
		X: util.HP_CORNERX,
		Y: util.HP_CORNERY,
	})
	HP.Push(pixel.Vec{
		X: (util.BAR_WIDTH * (m.PlayerHP / util.MAX_HP)) + util.HP_CORNERX,
		Y: util.BAR_HEIGHT + util.HP_CORNERY,
	})
	HP.Rectangle(0)

	//Draw Timer
	timer := imdraw.New(nil)
	timer.Color = pixel.RGB(.5, 0, 1)
	timer.Push(pixel.Vec{
		X: util.TIMER_CORNERX,
		Y: util.TIMER_CORNERY,
	})
	timer.Push(pixel.Vec{
		X: util.BAR_HEIGHT + util.TIMER_CORNERX,
		Y: (util.BAR_WIDTH * m.AttackTimer) + util.TIMER_CORNERY,
	})
	timer.Rectangle(0)

	cutter := m.Cutter.GetImage()
	attack := m.CurrentAttack.GetImage()

	//Write changes to canvas
	bbox.Draw(win)
	eHP.Draw(win)
	HP.Draw(win)
	timer.Draw(win)
	attack.Draw(win)
	cutter.Draw(win)

}
