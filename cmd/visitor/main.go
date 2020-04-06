package main

import (
	"fmt"

	"github.com/0ndreu/patterns-wb/pkg/visitor/areacalculator"
	"github.com/0ndreu/patterns-wb/pkg/visitor/circle"
	"github.com/0ndreu/patterns-wb/pkg/visitor/rectangle"
	"github.com/0ndreu/patterns-wb/pkg/visitor/visitor"
)

func main() {
	area := areacalculator.NewArea()
	circle := circle.NewCircle(area, 3)
	vis := visitor.NewVisitor(circle)
	a := vis.Accept()
	fmt.Println(a)
	rect := rectangle.NewCircle(area, 8, 10)
	vis = visitor.NewVisitor(rect)
	res := vis.Accept()
	fmt.Println(res)

}
