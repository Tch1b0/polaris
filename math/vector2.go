package math

import (
	"fmt"
	"math"
	"reflect"
)

// A vector with 2 dimensions (x, y)
type Vector2[T Number] struct {
	X T
	Y T
}

func (v *Vector2[T]) Add(o *Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v *Vector2[T]) Sub(o *Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v *Vector2[T]) Mul(o *Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v *Vector2[T]) ScalarMul(fac T) Vector2[T] {
	return Vector2[T]{
		X: v.X * fac,
		Y: v.Y * fac,
	}
}

// Calculate the distance between this vector and another one using the pythagora theorem
func (v *Vector2[T]) DistanceTo(o *Vector2[T]) float64 {
	return math.Sqrt(math.Pow(float64(o.X-v.X), 2) + math.Pow(float64(o.Y-v.Y), 2))
}

// Calculate the distance between this vector and another one using the manhattan distance
func (v *Vector2[T]) ManhattanDistanceTo(o *Vector2[T]) float64 {
	return math.Abs(float64(v.X-o.X)) + math.Abs(float64(v.Y-o.Y))
}

func (v *Vector2[T]) String() string {
	var x any = v.X
	var y any = v.Y

	numType := reflect.TypeOf(v.X).String()

	if numType == "float32" || numType == "float64" {
		return fmt.Sprintf("(%f|%f)", x, y)
	} else {
		return fmt.Sprintf("(%d|%d)", x, y)
	}
}
