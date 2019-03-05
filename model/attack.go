package model

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Attack interface {
	//Returns the image to draw this attack
	GetImage() *imdraw.IMDraw
	//Changes the image if it has some effect as a function of time
	Update(timestep int)
	//Accepts a cut from the given cutter object, and returns the ratio of the area of the halves from the cut
	AcceptCut(Cutter) float64
}

type staticCircleAttack struct {
	boundingBox pixel.Vec
	center      pixel.Vec
	rad         float64
}

func (s staticCircleAttack) GetImage() *imdraw.IMDraw {
	image := imdraw.New(nil)
	image.Color = pixel.RGB(1, 0, 0)
	adjustedCenter := pixel.Vec{
		X: s.center.X + s.boundingBox.X,
		Y: s.center.Y + s.boundingBox.Y,
	}
	image.Push(adjustedCenter)
	image.Circle(s.rad, 0)
	return image
}

func (s staticCircleAttack) Update(timestep int) {
	//Do not do anything, this is a static attack
}

func (s staticCircleAttack) AcceptCut(c Cutter) float64 {
	//TODO
	return 0
}

func NewStaticCircleAttack(center, boundingBox pixel.Vec, radius float64) Attack {
	return staticCircleAttack{
		center:      center,
		boundingBox: boundingBox,
		rad:         radius,
	}
}
