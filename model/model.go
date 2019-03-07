package model

import (
	"math"
	"math/rand"

	"github.com/DoubleWB/game_demo/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const CUT_SPAM_BUFFER = 30

type Model struct {
	CurrentAttack  Attack
	Cutter         Cutter
	BoundingCorner pixel.Vec
	EnemyHP        float64
	PlayerHP       float64
	AttackTimer    float64
	LastError      float64
	cutterAngle    float64
	cutSpamCount   int
}

func (m *Model) Update(win *pixelgl.Window) bool {
	//Handle Cutter Changes
	cdx := 0.0
	cdy := 0.0
	cdRot := 0.0
	currentPos := m.Cutter.GetPos()
	if win.Pressed(pixelgl.KeyLeft) && (currentPos.X-util.CUTTER_SPEED > 0) {
		cdx -= util.CUTTER_SPEED
	}

	if win.Pressed(pixelgl.KeyRight) && (currentPos.X+util.CUTTER_SPEED < util.BBOX_DIM) {
		cdx += util.CUTTER_SPEED
	}

	if win.Pressed(pixelgl.KeyDown) && (currentPos.Y-util.CUTTER_SPEED > 0) {
		cdy -= util.CUTTER_SPEED
	}

	if win.Pressed(pixelgl.KeyUp) && (currentPos.Y+util.CUTTER_SPEED < util.BBOX_DIM) {
		cdy += util.CUTTER_SPEED
	}

	if win.Pressed(pixelgl.KeyZ) {
		cdRot += util.ROTATE_SPEED
	}

	if win.Pressed(pixelgl.KeyX) {
		cdRot -= util.ROTATE_SPEED
	}

	m.cutterAngle += cdRot

	m.Cutter.Update(cdx, cdy, math.Tan(m.cutterAngle))

	m.AttackTimer -= .008
	if m.AttackTimer <= 0 || win.Pressed(pixelgl.KeySpace) && m.cutSpamCount == CUT_SPAM_BUFFER {
		//result is bounded by .5 and 0, .5 being the worst, and zero being the best
		result := m.CurrentAttack.AcceptCut(m.Cutter)
		m.LastError = result
		enemyDamage := 300 * (.5 - result)
		playerDamage := 300 * result
		m.EnemyHP -= enemyDamage
		m.PlayerHP -= playerDamage
		m.cutSpamCount = 0

		m.AttackTimer = 1.0
		randRad := rand.Float64() * 40
		randX := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad
		randY := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad
		m.CurrentAttack = NewStaticCircleAttack(m.BoundingCorner, pixel.Vec{X: randX, Y: randY}, randRad)

	} else if m.cutSpamCount < CUT_SPAM_BUFFER {
		m.cutSpamCount++
	}
	if m.PlayerHP <= 0.0 || m.EnemyHP <= 0.0 {
		return false
	}
	return true
}

func RestartModel() *Model {
	boundingBox := pixel.Vec{
		X: util.BBOX_CORNERX,
		Y: util.BBOX_CORNERY,
	}
	return &Model{
		CurrentAttack:  NewStaticCircleAttack(boundingBox, boundingBox, 40),
		Cutter:         NewBasicCutter(boundingBox),
		BoundingCorner: boundingBox,
		EnemyHP:        900.0,
		PlayerHP:       1000.0,
		AttackTimer:    1.0,
		cutterAngle:    0.0,
		cutSpamCount:   CUT_SPAM_BUFFER,
	}
}
