package model

import (
	"math"

	"github.com/DoubleWB/game_demo/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const HALF1 = false
const HALF2 = true

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
	fakeDraw := pixelgl.NewCanvas(pixel.R(0, 0, 1000, 1000))
	s.GetImage().Draw(fakeDraw)
	c.GetImage(false).Draw(fakeDraw)

	halves := make(map[bool]int)

	foundLine := false

	for row := util.BBOX_CORNERY; row < util.BBOX_CORNERY+util.BBOX_DIM; row += 1 {
		curHalf := HALF1
		pixCount := 0
		for x := util.BBOX_CORNERX; x < util.BBOX_CORNERX+util.BBOX_DIM; x += 1 {

			pix := fakeDraw.Color(pixel.Vec{X: x, Y: row})

			if pix == pixel.RGB(1.0, 0.0, 0.0) {
				pixCount += 1
			}

			if pix == pixel.RGB(0.0, 0.0, 0.0) {
				halves[curHalf] += pixCount
				pixCount = 0
				curHalf = HALF2
				foundLine = true
			}
		}

		//We didn't hit the line on this iteration
		if curHalf == HALF1 {
			if foundLine {
				halves[HALF2] += pixCount
			} else {
				halves[HALF1] += pixCount
			}
		} else {
			halves[HALF2] += pixCount
		}
	}

	sum := halves[HALF1] + halves[HALF2]
	c.PerformCut()

	return math.Abs(float64(halves[HALF1])/float64(sum) - .5)
}

func NewStaticCircleAttack(center, boundingBox pixel.Vec, radius float64) Attack {
	return staticCircleAttack{
		center:      center,
		boundingBox: boundingBox,
		rad:         radius,
	}
}
