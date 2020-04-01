package areacalculator

import (
	"math"
)

//
type Area interface {
	AreaForCircle(radius float64) (res float64)
	AreaForRectangle(length float64, height float64) (res float64)
}

type areaCalculator struct{}

//
func (a *areaCalculator) AreaForCircle(radius float64) (res float64) {
	res = math.Pow(radius, 2) * math.Pi
	return
}

//
func (a *areaCalculator) AreaForRectangle(length float64, height float64) (res float64) {
	res = length * height
	res = math.Abs(res)
	return
}

//
func NewArea() Area {
	return &areaCalculator{}
}
