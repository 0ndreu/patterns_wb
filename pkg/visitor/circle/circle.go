package circle

import "fmt"

type area interface {
	AreaForCircle(radius float64) (res float64)
}

type Circle interface {
	Accept() float64
}

type circle struct {
	area
	radius float64
}

// Accept return area of the circle
func (c *circle) Accept() (res float64) {
	res = c.area.AreaForCircle(c.radius)
	return
}

// NewCircle instance
func NewCircle(area area, radius float64) (c Circle, err error) {
	if radius > 0 {
		c = &circle{
			area:   area,
			radius: radius,
		}
		return
	}
	err = fmt.Errorf("there is no circle")
	c = nil
	return
}
