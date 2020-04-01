package circle

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

func (c *circle) Accept() (res float64) {
	res = c.area.AreaForCircle(c.radius)
	return
}

// NewCircle instance
func NewCircle(area area, radius float64) Circle {
	return &circle{
		area:   area,
		radius: radius,
	}
}
