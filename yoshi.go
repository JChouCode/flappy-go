package yoshi

const gravity = 1
const velocity = 0

type yoshi struct {
	X int
	Y int
}

func (y yoshi) update() {
	velocity += gravity
	y += velocity
}
