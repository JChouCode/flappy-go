package bird

import (
	"github.com/faiface/pixel"
	_ "image/png"
)

const gravity = 6
const airRes = 0.99

const lift = 300.0

type Bird struct {
	body pixel.Rect
	vel  pixel.Vec
}

func New() Bird {
	return Bird{pixel.R(500, 500, 570, 560), pixel.V(0, 0)}
}

func (b *Bird) Update(dt float64) {
	b.vel.Y -= gravity
	b.vel.Y *= airRes
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
	b.body = pixel.R(500, 500, 570, 560)
	b.vel = pixel.V(0, 0)
}
