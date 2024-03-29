package math

import (
	"fmt"
	"math"
)

// A vector with 2 dimensions (x, y)
type Vector2[T Number] struct {
	X T
	Y T
}

func (v Vector2[T]) Add(o Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vector2[T]) Sub(o Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vector2[T]) Mul(o Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v Vector2[T]) ScalarMul(fac T) Vector2[T] {
	return Vector2[T]{
		X: v.X * fac,
		Y: v.Y * fac,
	}
}

func (v Vector2[T]) ScalarMulVec(o Vector2[T]) T {
	return v.X*o.X + v.Y*o.Y
}

func (v Vector2[T]) Abs() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// alias for v.Abs()
func (v Vector2[T]) Length() T {
	return v.Abs()
}

func (v Vector2[T]) Normalize() Vector2[T] {
	l := v.Abs()
	return Vector2[T]{
		X: v.X / l,
		Y: v.Y / l,
	}
}

// Calculate the distance between this vector and another one using the pythagora theorem
func (v Vector2[T]) DistanceTo(o Vector2[T]) float64 {
	return math.Sqrt(math.Pow(float64(o.X-v.X), 2) + math.Pow(float64(o.Y-v.Y), 2))
}

// Calculate the distance between this vector and another one using the manhattan distance
func (v Vector2[T]) ManhattanDistanceTo(o Vector2[T]) float64 {
	return math.Abs(float64(v.X-o.X)) + math.Abs(float64(v.Y-o.Y))
}

func (v Vector2[T]) Array() [2]T {
	return [2]T{v.X, v.Y}
}

func (v Vector2[T]) String() string {
	return fmt.Sprintf("(%v|%v)", v.X, v.Y)
}
