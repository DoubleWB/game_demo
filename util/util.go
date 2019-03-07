package util

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

//Bounding Box
const BBOX_DIM = 500.0
const BBOX_CORNERX = 250.0
const BBOX_CORNERY = 250.0

//Enemy HP Bar
const EHP_CORNERX = 100.0
const EHP_CORNERY = 800.0
const MAX_HP = 1000.0

var HEALTH_COLOR = pixel.RGB(0.0, 1.0, .25)
var HEALTH_LOST_COLOR = pixel.RGB(1.0, 0.0, 0.0)

//Percent Error
const PERCENT_CORNERX = 600.0
const PERCENT_CORNERY = 800.0

//Player HP Bar
const HP_CORNERX = 600.0
const HP_CORNERY = 100.0

//Time Bar
const TIMER_CORNERX = 100.0
const TIMER_CORNERY = 250.0
const TIMER_RATE = .008

var TIMER_COLOR = pixel.RGB(.5, 0, 1)

//Bar General
const BAR_WIDTH = 200.0
const BAR_HEIGHT = 15.0

//Labels
const ENEMY_LABEL_X = 100.0
const ENEMY_LABEL_Y = 850.0
const PLAYER_LABEL_X = 600.0
const PLAYER_LABEL_Y = 150.0
const END_LABEL_X = 350.0
const END_LABEL_Y = 500.0

//Cutter
const CUT_SPAM_BUFFER = 30
const CUTTER_SPEED = 5.0
const ROTATE_SPEED = .05
const ANIM_FRAMES = 10
const CUTTER_THICK = 1.0
const CUTTER_IMAGE_THICK = 2.0
const CUTTER_ANIM_THICK = 5.0
const CUTTER_CENTER_THICK = 5.0

var CUTTER_ANIM_COLOR = pixel.RGB(1.0, 1.0, 0)
var CUTTER_COLOR = pixel.RGB(0, 0, 0)
var CUTTER_CENTER_COLOR = pixel.RGB(0, .5, .5)

//Attack
const RAD_LOWER_LIM = 10.0
const RAD_UPPER_LIM = 60.0
const ATTACK_DAMAGE = 100.0
const COLLISION_DAMAGE = 3.0

var ATTACK_COLOR = pixel.RGB(1.0, 0.0, 0.0)
var POISON_ATTACK_COLOR = pixel.RGB(1.0, 0, 1.0)

//MISC
const RATIO_CAP = .5

var BACKGROUND = colornames.Aliceblue
