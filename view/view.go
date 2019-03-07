package view

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"github.com/DoubleWB/game_demo/model"
	"github.com/DoubleWB/game_demo/util"
)

type View struct{}

//Get Bounding Box Image
func getBbox() *imdraw.IMDraw {
	bbox := imdraw.New(nil)
	bbox.Color = util.CUTTER_COLOR
	bbox.Push(pixel.Vec{
		X: util.BBOX_CORNERX,
		Y: util.BBOX_CORNERY,
	})
	bbox.Push(pixel.Vec{
		X: util.BBOX_DIM + util.BBOX_CORNERX,
		Y: util.BBOX_DIM + util.BBOX_CORNERY,
	})
	bbox.Rectangle(util.CUTTER_IMAGE_THICK)
	return bbox
}

//Get Enemy Health Points Image
func getEHP(m model.Model) *imdraw.IMDraw {
	eHP := imdraw.New(nil)
	eHP.Color = util.HEALTH_LOST_COLOR
	eHP.Push(pixel.Vec{
		X: util.EHP_CORNERX,
		Y: util.EHP_CORNERY,
	})
	eHP.Push(pixel.Vec{
		X: util.BAR_WIDTH + util.EHP_CORNERX,
		Y: util.BAR_HEIGHT + util.EHP_CORNERY,
	})
	eHP.Rectangle(0)
	eHP.Color = util.HEALTH_COLOR
	eHP.Push(pixel.Vec{
		X: util.EHP_CORNERX,
		Y: util.EHP_CORNERY,
	})
	eHP.Push(pixel.Vec{
		X: (util.BAR_WIDTH * (m.EnemyHP / util.MAX_HP)) + util.EHP_CORNERX,
		Y: util.BAR_HEIGHT + util.EHP_CORNERY,
	})
	eHP.Rectangle(0)
	return eHP
}

func getHP(m model.Model) *imdraw.IMDraw {
	HP := imdraw.New(nil)
	HP.Color = util.HEALTH_LOST_COLOR
	HP.Push(pixel.Vec{
		X: util.HP_CORNERX,
		Y: util.HP_CORNERY,
	})
	HP.Push(pixel.Vec{
		X: util.BAR_WIDTH + util.HP_CORNERX,
		Y: util.BAR_HEIGHT + util.HP_CORNERY,
	})
	HP.Rectangle(0)
	HP.Color = util.HEALTH_COLOR
	HP.Push(pixel.Vec{
		X: util.HP_CORNERX,
		Y: util.HP_CORNERY,
	})
	HP.Push(pixel.Vec{
		X: (util.BAR_WIDTH * (m.PlayerHP / util.MAX_HP)) + util.HP_CORNERX,
		Y: util.BAR_HEIGHT + util.HP_CORNERY,
	})
	HP.Rectangle(0)
	return HP
}

func getTimer(m model.Model) *imdraw.IMDraw {
	timer := imdraw.New(nil)
	timer.Color = util.TIMER_COLOR
	timer.Push(pixel.Vec{
		X: util.TIMER_CORNERX,
		Y: util.TIMER_CORNERY,
	})
	timer.Push(pixel.Vec{
		X: util.BAR_HEIGHT + util.TIMER_CORNERX,
		Y: (util.BAR_WIDTH * m.AttackTimer) + util.TIMER_CORNERY,
	})
	timer.Rectangle(0)
	return timer
}

func getPercent(m model.Model) *text.Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	percent := text.New(pixel.V(util.PERCENT_CORNERX, util.PERCENT_CORNERY), atlas)
	percent.Color = colornames.Black
	fmt.Fprintln(percent, fmt.Sprintf("Error: %.2f", m.LastError))
	return percent
}

func getEnemyLabel() *text.Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	label := text.New(pixel.V(util.ENEMY_LABEL_X, util.ENEMY_LABEL_Y), atlas)
	label.Color = colornames.Black
	fmt.Fprintln(label, "ENEMY:")
	return label
}

func getYouLabel() *text.Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	label := text.New(pixel.V(util.PLAYER_LABEL_X, util.PLAYER_LABEL_Y), atlas)
	label.Color = colornames.Black
	fmt.Fprintln(label, "YOU:")
	return label
}

func getWonLabel() *text.Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	label := text.New(pixel.V(util.END_LABEL_X, util.END_LABEL_Y), atlas)
	label.Color = util.HEALTH_COLOR
	fmt.Fprintln(label, "You WIN!!!!")
	return label
}

func getLostLabel() *text.Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	label := text.New(pixel.V(util.END_LABEL_X, util.END_LABEL_Y), atlas)
	label.Color = util.HEALTH_LOST_COLOR
	fmt.Fprintln(label, "You LOSE!!!")
	return label
}

func (v View) DrawToWindow(win *pixelgl.Window, m model.Model) {

	if m.HasLost || m.HasWon {
		if m.HasLost {
			label := getLostLabel()
			label.Draw(win, pixel.IM.Scaled(label.Orig, 5))
		} else if m.HasWon {
			label := getWonLabel()
			label.Draw(win, pixel.IM.Scaled(label.Orig, 5))
		}
		return
	}

	//draw bounding box
	bbox := getBbox()

	//draw enemy hp
	eHP := getEHP(m)

	//draw player hp
	HP := getHP(m)

	//Draw Timer
	timer := getTimer(m)

	//Draw percent
	percent := getPercent(m)

	//Draw Player Label
	youLabel := getYouLabel()

	//Draw Player Label
	enemyLabel := getEnemyLabel()

	//Draw cutter
	cutter := m.Cutter.GetImage(true, m.BoundingCorner)

	//Draw attack
	attack := m.CurrentAttack.GetImage(m.BoundingCorner)

	//Write changes to canvas
	//Order is important, last images are drawn over previous ones
	bbox.Draw(win)
	eHP.Draw(win)
	HP.Draw(win)
	percent.Draw(win, pixel.IM.Scaled(percent.Orig, 4))
	youLabel.Draw(win, pixel.IM.Scaled(youLabel.Orig, 4))
	enemyLabel.Draw(win, pixel.IM.Scaled(enemyLabel.Orig, 4))
	timer.Draw(win)
	attack.Draw(win)
	cutter.Draw(win)

}
