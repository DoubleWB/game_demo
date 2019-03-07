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
	LastError      float64
	HasWon         bool
	HasLost        bool
	cutterAngle    float64
	cutSpamCount   int
}

func (m *Model) updateCutter(win *pixelgl.Window) {
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
}

func (m Model) performCut() float64 {
	//Draw the images to a fake canvas to calculate the division graphically
	fakeDraw := pixelgl.NewCanvas(pixel.R(0, 0, util.BBOX_DIM, util.BBOX_DIM))
	fakeDraw.Clear(util.BACKGROUND)
	m.CurrentAttack.GetImage(pixel.V(0, 0)).Draw(fakeDraw)
	m.Cutter.GetImage(false, pixel.V(0, 0)).Draw(fakeDraw)

	halves := make(map[bool]int)

	//Once we have found the line, any rows where we don't see it at all belong to half 2
	foundLine := false

	for row := 0.0; row < util.BBOX_DIM; row += 1 {
		curHalf := HALF1
		pixCount := 0
		for x := 0.0; x < util.BBOX_DIM; x += 1 {

			pix := fakeDraw.Color(pixel.Vec{X: x, Y: row})

			if pix != pixel.ToRGBA(util.BACKGROUND) {
				//if the pixel is not the background color, and is not the cutter color, it is part of an attack, so add it to this "batch"
				//if the pixel is part of the cutter color, you add the whole batch to the half you are currently on, and switch to the next half
				//since the pixels after the line will be a part of another half
				if pix == util.CUTTER_COLOR {
					halves[curHalf] += pixCount
					pixCount = 0
					curHalf = HALF2
					foundLine = true
				} else {
					pixCount += 1
				}
			}
		}

		//If we never switched halves on this row
		if curHalf == HALF1 {
			//and we haven't EVER seen the line, this batch belongs to the first half
			//and if we have seen the line, this batch belongs to the second half
			if foundLine {
				halves[HALF2] += pixCount
			} else {
				halves[HALF1] += pixCount
			}
		} else {
			//otherwise, all leftover baches belong to the second half
			halves[HALF2] += pixCount
		}
	}

	sum := halves[HALF1] + halves[HALF2]

	//Have the cutter perform it's animation or whatever other changes it wants to
	m.Cutter.PerformCut()

	//absolute distance from a half ratio of 1:1
	return math.Abs(float64(halves[HALF1])/float64(sum) - util.RATIO_CAP)
}

//Handle the logic for calculating damage after a cut, and generating the next attack
func (m *Model) handleCut() {
	//result is bounded by .5 and 0, .5 being the worst, and zero being the best
	result := m.performCut()
	m.LastError = result
	//enemy takes damage capping at result = 0 (best)
	enemyDamage := util.ATTACK_DAMAGE * (util.RATIO_CAP - result)
	//player takes damagee capping at result = .5 (worst)
	playerDamage := util.ATTACK_DAMAGE * result
	m.EnemyHP -= enemyDamage
	m.PlayerHP -= playerDamage
	m.cutSpamCount = 0

	//prepare at least one set of random values for a circle attack
	m.AttackTimer = 1.0
	randRad := util.RAD_LOWER_LIM + (rand.Float64() * util.RAD_UPPER_LIM)
	randX := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad
	randY := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad

	//select a random new attack
	attackSelector := rand.Float64()
	if attackSelector <= .33 {
		//random basic attack
		m.CurrentAttack = NewStaticCircleAttack(pixel.Vec{X: randX, Y: randY}, randRad)
	} else if attackSelector <= .66 {
		//random poison attack
		m.CurrentAttack = NewPoisonCircleAttack(pixel.Vec{X: randX, Y: randY}, randRad)
	} else {
		//random multi attack
		//calculate another attack to add to the multiattack
		attacks := []Attack{NewStaticCircleAttack(pixel.Vec{X: randX, Y: randY}, randRad)}
		randRad := util.RAD_LOWER_LIM + (rand.Float64() * util.RAD_UPPER_LIM)
		randX := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad
		randY := (rand.Float64() * (util.BBOX_DIM - (2 * randRad))) + randRad
		attacks = append(attacks, NewPoisonCircleAttack(pixel.Vec{X: randX, Y: randY}, randRad))
		m.CurrentAttack = NewMultiAttack(attacks)
	}
}

//If the player is in collision with an attack, take collision damage
func (m *Model) handleCollision() {
	if m.CurrentAttack.InCollision(m.Cutter) {
		m.PlayerHP -= util.COLLISION_DAMAGE
	}
}

func (m *Model) Update(win *pixelgl.Window) {

	if !m.HasLost && !m.HasWon {

		m.updateCutter(win)

		m.handleCollision()

		m.AttackTimer -= util.TIMER_RATE
		//you can only cut once every CUT_SPAM_BUFFER frames
		if m.AttackTimer <= 0 || win.Pressed(pixelgl.KeySpace) && m.cutSpamCount == util.CUT_SPAM_BUFFER {
			m.handleCut()
		} else if m.cutSpamCount < util.CUT_SPAM_BUFFER {
			m.cutSpamCount++
		}
	}

	if m.PlayerHP <= 0.0 {
		m.HasLost = true
	} else if m.EnemyHP <= 0.0 {
		m.HasWon = true
	}
}

func RestartModel() *Model {
	boundingBox := pixel.Vec{
		X: util.BBOX_CORNERX,
		Y: util.BBOX_CORNERY,
	}
	return &Model{
		CurrentAttack:  NewStaticCircleAttack(boundingBox, 40),
		Cutter:         NewBasicCutter(boundingBox),
		BoundingCorner: boundingBox,
		EnemyHP:        util.MAX_HP,
		PlayerHP:       util.MAX_HP,
		AttackTimer:    1.0,
		cutterAngle:    0.0,
		cutSpamCount:   util.CUT_SPAM_BUFFER,
	}
}
