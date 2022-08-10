package main

import (
	"fmt"
	p "level_one/24/point"
)

/*
	Разработать программу нахождения расстояния между двумя точками, которые
	представлены в виде структуры Point с инкапсулированными параметрами x,y и
	конструктором.
*/

func main() {
	p1, p2 := p.NewPoint(10, 0), p.NewPoint(20, 10)

	d1 := p.CalculateDistance(p1, p2)
	d2 := p1.GetDistanceToPoint(p2)

	fmt.Printf("Distance between points P1 (%f, %f) and P2 (%f, %f) equals to %f\n", p1.GetX(), p1.GetY(), p2.GetX(), p2.GetY(), d1)
	fmt.Printf("Distance was calculated by two methods, both gave same result: %t\n", d1 == d2)
}
