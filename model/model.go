package model

import (
	"math"
	"math/rand"

	"github.com/DoubleWB/game_demo/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Model struct {
	CurrentAttack  Attack
	Cutter         Cutter
	BoundingCorner pixel.Vec
	EnemyHP        float64
	PlayerHP       float64
	AttackTimer    float64
	cutterAngle    float64
}

func (m *Model) Update(win *pixelgl.Window) {
	//Handle Cutter Changes
	cdx := 0.0
	cdy := 0.0
	cdRot := 0.0
	if win.Pressed(pixelgl.KeyLeft) {
		cdx -= util.CUTTER_SPEED
	}

	if win.Pressed(pixelgl.KeyRight) {
		cdx += util.CUTTER_SPEED
	}

	if win.Pressed(pixelgl.KeyDown) {
		cdy -= util.CUTTER_SPEED
	}

	if win.Pressed(pixelgl.KeyUp) {
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

	hasCut := false

	m.AttackTimer -= .008
	if m.AttackTimer <= 0 || win.Pressed(pixelgl.KeySpace) {
		hasCut = true
		m.AttackTimer = 1.0
		randRad := rand.Float64() * 40
		randX := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad
		randY := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad
		m.CurrentAttack = NewStaticCircleAttack(m.BoundingCorner, pixel.Vec{X: randX, Y: randY}, randRad)
	}

	if hasCut {
		result := m.CurrentAttack.AcceptCut(m.Cutter)
		enemyDamage := 50 * result
		playerDamage := 50 * (1.0 - result)
		m.EnemyHP -= enemyDamage
		m.PlayerHP -= playerDamage
	}
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
	}
}
