package model

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/DoubleWB/game_demo/util"
)

type Cutter interface {
	//Returns the image to draw this cutter
	GetImage() *imdraw.IMDraw
	//Sets the cutter paramenters
	Set(x, y, slope float64)
	//Adds the given delta values to the cutter parameters
	Update(x, y, slope float64)
}

type basicCutter struct {
	boundingBox pixel.Vec
	center      pixel.Vec
	slope       float64
}

func (b basicCutter) GetImage() *imdraw.IMDraw {
	image := imdraw.New(nil)
	image.Color = pixel.RGB(0, 0, 0)

	if b.slope != 0 {
		//Calculate and push whatever positive intercept is in one corner of the bounding box
		xInt := pixel.Vec{
			X: (-b.center.Y + (b.slope * b.center.X)) / b.slope,
			Y: 0,
		}
		yInt := pixel.Vec{
			X: 0,
			Y: (b.slope * (-b.center.X)) + b.center.Y,
		}

		if xInt.X >= 0 {
			image.Push(pixel.Vec{
				X: b.boundingBox.X + xInt.X,
				Y: b.boundingBox.Y,
			})
		} else {
			image.Push(pixel.Vec{
				X: b.boundingBox.X,
				Y: b.boundingBox.Y + yInt.Y,
			})
		}

		//Calculate and push whatever positive intercept is in the other corner of the bounding box
		xUpperInt := pixel.Vec{
			X: ((util.BBOX_DIM - b.center.Y) / b.slope) + b.center.X,
			Y: util.BBOX_DIM,
		}
		yUpperInt := pixel.Vec{
			X: util.BBOX_DIM,
			Y: (b.slope * (util.BBOX_DIM - b.center.X)) + b.center.Y,
		}

		if xInt.X <= util.BBOX_DIM {
			image.Push(pixel.Vec{
				X: b.boundingBox.X + xUpperInt.X,
				Y: b.boundingBox.Y + util.BBOX_DIM,
			})
		} else {
			image.Push(pixel.Vec{
				X: b.boundingBox.X + util.BBOX_DIM,
				Y: b.boundingBox.Y + yUpperInt.Y,
			})
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
	image.Line(2.0)

	adjustedCenter := pixel.Vec{
		X: b.center.X + b.boundingBox.X,
		Y: b.center.Y + b.boundingBox.Y,
	}
	image.Color = pixel.RGB(0, .5, .5)
	image.Push(adjustedCenter)
	image.Circle(5.0, 0)
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
