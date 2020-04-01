package visitor

type Area interface {
	Accept() float64
}

type Visitor interface {
	Accept() float64
}

type visitor struct{
	area Area
}

//
func (v *visitor) Add (a Area) {
	v.area = a
}

//
func (v *visitor) Accept() (res float64) {
	res = v.area.Accept()
	return

}

//
func NewVisitor(c Area) Visitor {
	return &visitor{area:c}
}
