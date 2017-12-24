package main

import "fmt"

type shape interface {
    getArea() float64
}

type square struct {
    sideLength float64
}
type triangle struct {
    height float64
    base float64
}

func main() {
    s := square { sideLength: 10 }
    t := triangle { height: 11, base: 10 }

    printArea(s)
    printArea(t)
}

func printArea(s shape) {
    area := s.getArea()
    fmt.Println("Area:", area)
}

func (s square) getArea() float64 {
    area := s.sideLength * s.sideLength
    return area
}

func (t triangle) getArea() float64 {
    area := 0.5 * t.base * t.height
    return area
}
