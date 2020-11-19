package day2

import "testing"

import "math"

import "fmt"

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func Test1(t *testing.T) {
	f := MyFloat(math.Sqrt2)
	t.Log(f.Abs())
}

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Test2(t *testing.T) {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}

	a = f
	// a = &v

	fmt.Println(a.Abs())
}
