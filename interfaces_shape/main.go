package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct{}
type square struct{}

func main() {

	t := triangle{}
	s := square{}

	printArea(t)
	printArea(s)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (triangle) getArea() float64 {
	return 2.0
}

func (square) getArea() float64 {
	return 3.1
}
