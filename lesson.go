package main

import "fmt"

type (v Vertex) Area() int{
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex3D struct{
	Vertex
	z int
}

func (v Vertex3D) Area3D() int{
	return v.x * v.y
}

func (v Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y int) *Vertex3D{
	return &Vertex3D{x, y}
}

func main() {
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}