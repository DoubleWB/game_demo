package model

import (
	"github.com/DoubleWB/game_demo/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const HALF1 = false
const HALF2 = true

type Attack interface {
	//Returns the image to draw this attack
	GetImage(boundingBox pixel.Vec) *imdraw.IMDraw
	//Changes the image if it has some effect as a function of time
	Update(timestep int)
	//Returns true if this attack is in collision with the given cutter
	InCollision(Cutter) bool
}

type staticCircleAttack struct {
	center pixel.Vec
	rad    float64
}

func (s staticCircleAttack) GetImage(boundingBox pixel.Vec) *imdraw.IMDraw {
	image := imdraw.New(nil)
	image.Color = util.ATTACK_COLOR
	adjustedCenter := pixel.Vec{
		X: s.center.X + boundingBox.X,
		Y: s.center.Y + boundingBox.Y,
	}
	image.Push(adjustedCenter)
	image.Circle(s.rad, 0)
	return image
}

func (s staticCircleAttack) Update(timestep int) {
	//Do not do anything, this is a static attack
}

//This attack is never in collision
func (s staticCircleAttack) InCollision(c Cutter) bool {
	return false
}

func NewStaticCircleAttack(center pixel.Vec, radius float64) Attack {
	return staticCircleAttack{
		center: center,
		rad:    radius,
	}
}

//Attack that also does collision damage
type poisonCircleAttack struct {
	center pixel.Vec
	rad    float64
}

func (p poisonCircleAttack) GetImage(boundingBox pixel.Vec) *imdraw.IMDraw {
	image := imdraw.New(nil)
	image.Color = util.POISON_ATTACK_COLOR
	adjustedCenter := pixel.Vec{
		X: p.center.X + boundingBox.X,
		Y: p.center.Y + boundingBox.Y,
	}
	image.Push(adjustedCenter)
	image.Circle(p.rad, 0)
	return image
}

func (p poisonCircleAttack) Update(timestep int) {
	//Do not do anything, this is a static attack
}

func (p poisonCircleAttack) InCollision(c Cutter) bool {
	distance := p.center.Sub(c.GetPos()).Len()
	return distance <= p.rad
}

func NewPoisonCircleAttack(center pixel.Vec, radius float64) Attack {
	return poisonCircleAttack{
		center: center,
		rad:    radius,
	}
}

//Attack composed of an arbitrary number of other attacks
type multiAttack struct {
	attacks []Attack
}

func (m multiAttack) GetImage(boundingBox pixel.Vec) *imdraw.IMDraw {
	image := imdraw.New(nil)
	for _, a := range m.attacks {
		a.GetImage(boundingBox).Draw(image)
	}
	return image
}

func (m multiAttack) Update(timestep int) {
	//Do not do anything, this is a static attack
}

func (m multiAttack) InCollision(c Cutter) bool {
	inCollision := false
	for _, a := range m.attacks {
		inCollision = inCollision || a.InCollision(c)
	}
	return inCollision
}

func NewMultiAttack(attacks []Attack) Attack {
	return multiAttack{
		attacks: attacks,
	}
}
