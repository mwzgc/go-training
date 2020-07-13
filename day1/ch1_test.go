package day1

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestCh1(t *testing.T) {
	t.Log("hello world")
}

type T struct {
	name  string // name of the object
	value int    // its value
}

func (t *T) toString() {
	fmt.Println(fmt.Sprintf("name=%s, value=%d", t.name, t.value))
}

func add(x, y int) int {
	return x + y
}

func TestCh2(t *testing.T) {
	t.Log("The time is", time.Now())
	t.Log(add(1, 2))
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestVertex(t *testing.T) {
	v := Vertex{3, 4}
	t.Log(v.Abs())
}
