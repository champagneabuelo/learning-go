package main

import "fmt"

type Dump interface {
	Dump()
}

type Moto struct {
	Make  string
	Model string
	MPG   int
	Price float32
}

func (m Moto) Dump() {
	fmt.Println(m)
}

func (m *Moto) Discount() {
	m.Price = m.Price * 0.9
}

func x(dumps []Dump) {
	for _, d := range dumps {
		d.Dump()
	}
}

func main() {
	m1 := Moto{"BMW", "R100RT", 12, 14999.99}
	m2 := Moto{"Yamaha", "FZ600", 18, 12000.00}
	m3 := Moto{"Harley", "Soft Tail", 11, 15000.00}
	m1.Dump()
	m2.Dump()
	m3.Dump()
	m1.Discount()
	m1.Dump()
	dumps := []Dump{m1, m2, m3}
	x(dumps[:])

}
