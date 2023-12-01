package math

import (
	"fmt"
	"math"
)

// A vector with 3 dimensions (x, y, z)
type Vector3[T Number] struct {
	X T
	Y T
	Z T
}

func (v Vector3[T]) Add(o Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v Vector3[T]) Sub(o Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func (v Vector3[T]) Mul(o Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
	}
}

func (v Vector3[T]) ScalarMul(fac T) Vector3[T] {
	return Vector3[T]{
		X: v.X * fac,
		Y: v.Y * fac,
		Z: v.Z * fac,
	}
}

func (v Vector3[T]) ScalarMulVec(o Vector3[T]) T {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector3[T]) Abs() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

// alias for v.Abs()
func (v Vector3[T]) Length() T {
	return v.Abs()
}

func (v Vector3[T]) Normalize() Vector3[T] {
	l := v.Abs()
	return Vector3[T]{
		X: v.X / l,
		Y: v.Y / l,
		Z: v.Z / l,
	}
}

func (v Vector3[T]) CrossProduct(o Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: v.Y*o.Z - v.Z*o.Y,
		Y: v.Z*o.X - v.X*o.Z,
		Z: v.X*o.Y - v.Y*o.X,
	}
}

// Calculate the distance between this vector and another one using the pythagora theorem
func (v Vector3[T]) DistanceTo(o Vector3[T]) float64 {
	return math.Sqrt(
		math.Pow(math.Sqrt(
			math.Pow(float64(o.X-v.X), 2)+
				math.Pow(float64(o.Y-v.Y), 2),
		), 2) + math.Pow(float64(o.Z-v.Z), 2),
	)
}

// Calculate the distance between this vector and another one using the manhattan distance
func (v Vector3[T]) ManhattanDistanceTo(o Vector3[T]) float64 {
	return math.Abs(float64(v.X-o.X)) + math.Abs(float64(v.Y-o.Y)) + math.Abs(float64(v.Z-o.Z))
}

func (v Vector3[T]) Array() [3]T {
	return [3]T{v.X, v.Y, v.Z}
}

func (v Vector3[T]) String() string {
	return fmt.Sprintf("(%v|%v|%v)", v.X, v.Y, v.Z)
}
