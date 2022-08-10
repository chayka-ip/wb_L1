package point

import "math"

/*
	Coordinates are inacessible outside the point package
*/
type Point struct {
	x, y float64
}

// public constructor
func NewPoint(x, y float64) *Point {
	return &Point{
		x, y,
	}
}

// coordinate getter
func (p *Point) GetX() float64 {
	return p.x
}

// coordinate getter
func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) GetDistanceToPoint(other *Point) float64 {
	return CalculateDistance(p, other)
}

func CalculateDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
