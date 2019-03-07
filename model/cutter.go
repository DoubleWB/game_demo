package model

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/DoubleWB/game_demo/util"
)

type Cutter interface {
	//Returns the image to draw this cutter
	GetImage(withCenter bool) *imdraw.IMDraw
	//Sets the cutter paramenters
	Set(x, y, slope float64)
	//Adds the given delta values to the cutter parameters
	Update(x, y, slope float64)
	//Returns the position of this cutter
	GetPos() pixel.Vec
	//Performs Cut
	PerformCut()
}

type basicCutter struct {
	boundingBox pixel.Vec
	center      pixel.Vec
	slope       float64
	cutAnim     int
}

func (b *basicCutter) GetImage(withCenter bool) *imdraw.IMDraw {
	image := imdraw.New(nil)
	if b.cutAnim > 0 && withCenter {
		image.Color = pixel.RGB(1.0, 1.0, 0)
		b.cutAnim--
	} else {
		image.Color = pixel.RGB(0, 0, 0)
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

		intercepts := []pixel.Vec{xInt, yInt, xUpperInt, yUpperInt}

		for _, intercept := range intercepts {
			if intercept.X <= util.BBOX_DIM && intercept.X >= 0 && intercept.Y <= util.BBOX_DIM && intercept.Y >= 0 {
				image.Push(pixel.Vec{
					X: b.boundingBox.X + intercept.X,
					Y: b.boundingBox.Y + intercept.Y,
				})
			}
		}
	} else {
		image.Push((pixel.Vec{
			X: b.boundingBox.X,
			Y: b.boundingBox.Y + b.center.Y,
		}))
		image.Push((pixel.Vec{
			X: b.boundingBox.X + util.BBOX_DIM,
			Y: b.boundingBox.Y + b.center.Y,
		}))
	}

	//Draw cutter line
	if b.cutAnim > 0 && withCenter {
		image.Line(5.0)
	} else {
		image.Line(2.0)
	}

	if withCenter {
		adjustedCenter := pixel.Vec{
			X: b.center.X + b.boundingBox.X,
			Y: b.center.Y + b.boundingBox.Y,
		}
		image.Color = pixel.RGB(0, .5, .5)
		image.Push(adjustedCenter)
		image.Circle(5.0, 0)
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

func (b *basicCutter) PerformCut() {
	b.cutAnim = 10
}

func NewBasicCutter(boundingBox pixel.Vec) Cutter {
	return &basicCutter{
		center: pixel.Vec{
			X: util.BBOX_DIM / 2,
			Y: util.BBOX_DIM / 2,
		},
		boundingBox: boundingBox,
		slope:       0,
	}
}
