package areacalculator

import (
	"math"
)

// Area give access to AreaForCircle and AreaForRectangle
type Area interface {
	AreaForCircle(radius float64) (res float64)
	AreaForRectangle(length float64, height float64) (res float64)
}

type areaCalculator struct{}

// AreaForCircle calculate area of circle
func (a *areaCalculator) AreaForCircle(radius float64) (res float64) {
	res = math.Pow(radius, 2) * math.Pi
	return
}

// AreaForRectangle calculate area of rectangle
func (a *areaCalculator) AreaForRectangle(length float64, height float64) (res float64) {
	res = length * height
	res = math.Abs(res)
	return
}

// NewAre instance
func NewArea() Area {
	return &areaCalculator{}
}
