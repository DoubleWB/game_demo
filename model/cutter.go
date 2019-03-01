package model

import (
	"github.com/faiface/pixel/imdraw"
)

type Cutter interface {
	//Returns the image to draw this attack
	GetImage() *imdraw.IMDraw
	//
	Update(timestep int)
}
