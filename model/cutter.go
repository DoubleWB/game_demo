package model

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/DoubleWB/game_demo/util"
)

type Cutter interface {
	//Returns the image to draw this cutter. Fancy indicates that this is for a graphic presentation, and not cutting calculation,
	//and boundingBox is the location of the bounding box of the play window, to adjust the drawing accordingly.
	GetImage(fancy bool, boundingBox pixel.Vec) *imdraw.IMDraw
	//Sets the cutter paramenters
	Set(x, y, slope float64)
	//Adds the given delta values to the cutter parameters
	Update(x, y, slope float64)
	//Returns the position of this cutter
	GetPos() pixel.Vec
	//Returns the slope of this cutter
	GetSlope() float64
	//Executes all actions associated with performing a cut
	PerformCut()
}

type basicCutter struct {
	center  pixel.Vec
	slope   float64
	cutAnim int
}

func (b *basicCutter) GetImage(fancy bool, boundingBox pixel.Vec) *imdraw.IMDraw {
	image := imdraw.New(nil)
	if b.cutAnim > 0 && fancy {
		image.Color = util.CUTTER_ANIM_COLOR
		b.cutAnim--
	} else {
		image.Color = util.CUTTER_COLOR
	}

	if b.slope != 0 {
		//Calculate intercepts with all four bounds
		xInt := pixel.Vec{
			X: (-b.center.Y + (b.slope * b.center.X)) / b.slope,
			Y: 0,
		}
		yInt := pixel.Vec{
			X: 0,
			Y: (b.slope * (-b.center.X)) + b.center.Y,
		}
		xUpperInt := pixel.Vec{
			X: ((util.BBOX_DIM - b.center.Y) / b.slope) + b.center.X,
			Y: util.BBOX_DIM,
		}
		yUpperInt := pixel.Vec{
			X: util.BBOX_DIM,
			Y: (b.slope * (util.BBOX_DIM - b.center.X)) + b.center.Y,
		}

		//Go through all of the intercepts - if any of them are actually outside the window, we don't want to deal with them.
		//The only case where this will not eliminate exactly 2 verts is a line that intersects the corners of the box,
		//in which case the two intercepts at each corner are the same, so we don't care about drawing both of them.
		intercepts := []pixel.Vec{xInt, yInt, xUpperInt, yUpperInt}
		for _, intercept := range intercepts {
			if intercept.X <= util.BBOX_DIM && intercept.X >= 0 && intercept.Y <= util.BBOX_DIM && intercept.Y >= 0 {
				image.Push(pixel.Vec{
					X: boundingBox.X + intercept.X,
					Y: boundingBox.Y + intercept.Y,
				})
			}
		}
	} else { // if the slope is zero, special case to avoid dividing by zero
		image.Push((pixel.Vec{
			X: boundingBox.X,
			Y: boundingBox.Y + b.center.Y,
		}))
		image.Push((pixel.Vec{
			X: boundingBox.X + util.BBOX_DIM,
			Y: boundingBox.Y + b.center.Y,
		}))
	}

	//Draw cutter line
	if b.cutAnim > 0 && fancy {
		image.Line(util.CUTTER_ANIM_THICK)
	} else if fancy {
		image.Line(util.CUTTER_IMAGE_THICK)
	} else {
		image.Line(util.CUTTER_THICK)
	}

	if fancy {
		adjustedCenter := pixel.Vec{
			X: b.center.X + boundingBox.X,
			Y: b.center.Y + boundingBox.Y,
		}
		image.Color = util.CUTTER_CENTER_COLOR
		image.Push(adjustedCenter)
		image.Circle(util.CUTTER_CENTER_THICK, 0)
	}

	return image
}

func (b *basicCutter) Set(x, y, slope float64) {
	b.center.X = x
	b.center.Y = y
	b.slope = slope
}

func (b *basicCutter) Update(dX, dY, newSlope float64) {
	b.center.X += dX
	b.center.Y += dY
	b.slope = newSlope
}

func (b basicCutter) GetPos() pixel.Vec {
	return b.center
}

func (b basicCutter) GetSlope() float64 {
	return b.slope
}

func (b *basicCutter) PerformCut() {
	b.cutAnim = util.ANIM_FRAMES
}

func NewBasicCutter(boundingBox pixel.Vec) Cutter {
	return &basicCutter{
		center: pixel.Vec{
			X: util.BBOX_DIM / 2,
			Y: util.BBOX_DIM / 2,
		},
		slope: 0,
	}
}
