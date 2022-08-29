package main

import (
	"fmt"
)

type Vec2 interface {
	XY() (int, int)
}

type Vector struct {
	x, y int
}

type Unvector struct {
	Vector
}

func(v Vector)XY() (int, int) {
	return v.x, v.y
}

func (v Vector)Sum(a Vec2) Vector {
	x1, y1 := v.XY()
	x2, y2 := a.XY()

	x1 += x2
	y1 += y2

	return Vector{x1, y1}
}

func NewUnvector(x, y int) Unvector {
	v := Unvector{Vector{x, y}}
	return v
}

func V(v Vec2) Unvector {
	x, y := v.XY()
	return NewUnvector(x, y)
}

func (v Unvector)Mul(m Vec2) Unvector {
	x1, y1 := v.XY()
	x2, y2 := m.XY()

	x1 *= x2
	y1 *= y2

	return NewUnvector(x1, y1)
}

func main() {
	u := NewUnvector(1, 1)
	u = V(u.Sum(Vector{2, 2}))
	u = V(u.Mul(NewUnvector(2, 2)))
	x, y := u.XY()
	fmt.Println(x, y)
}

