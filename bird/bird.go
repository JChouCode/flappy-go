package bird

import (
	"github.com/faiface/pixel"
	_ "image/png"
)

const sheight = 512.0
const swidth = 672.0
const bheight = 24.0
const bwidth = 24.0

const gravity = 10
const airRes = 0.99

const lift = 250.0

type Bird struct {
	body pixel.Rect
	vel  pixel.Vec
}

func New() Bird {
	minY := sheight/2 - bheight/2
	minX := swidth * 0.4
	return Bird{pixel.R(minX, minY, minX+bwidth, minX+bheight), pixel.V(0, 0)}
}

func (b *Bird) Update(dt float64) {
	b.vel.Y -= gravity
	b.vel.Y *= airRes
	if b.body.Center().Y <= 112 && b.vel.Y < 0 {
		return
	}
	b.body = b.body.Moved(b.vel.Scaled(dt))
}

func (b *Bird) Lift() {
	b.vel.Y = lift
}

func (b Bird) GetVel() pixel.Vec {
	return b.vel
}

func (b Bird) GetBody() pixel.Rect {
	return b.body
}

func GetLift() float64 {
	return lift
}

func (b *Bird) Reset() {
	minY := sheight/2 - bheight/2
	minX := swidth * 0.4
	b.body = pixel.R(minX, minY, minX+bwidth, minX+bheight)
	b.vel = pixel.V(0, 0)
}

func (b Bird) Falling() bool {
	return b.vel.Y < 0
}

func (b Bird) StopAnim() bool {
	return b.vel.Y < 6
}
