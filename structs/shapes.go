package shapes

import "math"

// Shape is implemented by anything that can tell us its Area
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	width  float64
	height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter returns the perimeter of the rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// Circle has the dimensions of a circle
type Circle struct {
	radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Triangle has the dimensions of a triangle
type Triangle struct {
	base   float64
	height float64
}

// Area returns the area of the triangle
func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}
