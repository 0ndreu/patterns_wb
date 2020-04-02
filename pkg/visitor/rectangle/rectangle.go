package rectangle

import "fmt"

type area interface {
	AreaForRectangle(length float64, height float64) (res float64)
}

type Rectangle interface {
	Accept() float64
}

type rectangle struct {
	area
	length float64
	height float64
}

// Accept return area of rectangle
func (c *rectangle) Accept() (res float64) {
	res = c.area.AreaForRectangle(c.length, c.height)
	return
}

// NewCircle instance
func NewRectangle(area area, length float64, height float64) (r Rectangle, err error) {
	if length > 0 && height > 0 {
		r = &rectangle{
			area:   area,
			length: length,
			height: height,
		}
		return
	}
	err = fmt.Errorf("there is no rectangle")
	r = nil
	return
}
