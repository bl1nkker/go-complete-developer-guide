package interfaces

import "fmt"

type shape interface{
	GetArea() float64
}

type Rectangle struct{
	W float64
	H float64
}
type Circle struct{
	D float64
}
type Triangle struct{
	K1 float64
	K2 float64
	H float64
}

func (r Rectangle) GetArea() float64{
	return r.W * r.H
}
func (t Triangle) GetArea() float64{
	return t.K1 + t.K2 + t.H
}
func (c Circle) GetArea() float64{
	return c.D * 3.14
}

func PrintArea(s shape){
	fmt.Println("The area is:", s.GetArea())
}

func RunAssignment1(){
	circle := Circle{D: 10.0}
	rectanlge := Rectangle{H: 5.0, W: 4.0}
	triangle := Triangle{K1: 3.0, K2: 5.5, H: 7.0}
	PrintArea(circle)
	PrintArea(rectanlge)
	PrintArea(triangle)
}