package visitor

type Area interface {
	Accept() float64
}

type Visitor interface {
	Accept() float64
}

type visitor struct {
	area Area
}

// Accept returns area of the shape
func (v *visitor) Accept() (res float64) {
	res = v.area.Accept()
	return
}

// NewVisitor instance
func NewVisitor(c Area) Visitor {
	return &visitor{area: c}
}
