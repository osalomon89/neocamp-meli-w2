package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type SouCirculo interface {
	EuSouUmCirculo()
}

type Circle struct {
	Radius float64
}

// Método que implementa la interfaz Shape para Circle
func (c Circle) Area() float64 {
	return math.Phi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) EuSouUmCirculo() {
	fmt.Println("Eu sou um circulo")
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Método que implementa la interfaz Shape para Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func SouUmCirculo(s Circle) {
	s.EuSouUmCirculo()
}

func main() {
	// areaRect := areaRectangle(15, 3)
	// areaCirc := areaCircle(3)
	// fmt.Println(areaRect)
	// fmt.Println(areaCirc)

	circle := Circle{Radius: 2.0}
	rectangle := Rectangle{Width: 3.0, Height: 4.0}

	PrintArea(circle)    // Polimorfismo con Circle
	PrintArea(rectangle) // Polimorfismo con Rectangle
	SouUmCirculo(circle)
}

func areaRectangle(width, height float64) float64 {
	return width * height
}

func areaCircle(radius float64) float64 {
	return math.Phi * radius * radius
}
