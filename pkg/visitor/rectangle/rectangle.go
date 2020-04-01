package rectangle

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

func (c *rectangle) Accept() (res float64) {
	res = c.area.AreaForRectangle(c.length, c.height)
	return
}

// NewCircle instance
func NewCircle(area area, length float64, height float64) Rectangle {
	return &rectangle{
		area:   area,
		length: length,
		height: height,
	}
}
