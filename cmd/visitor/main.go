package main

import (
	"fmt"
	"patterns-wb/pkg/visitor/areacalculator"
	"patterns-wb/pkg/visitor/circle"
	"patterns-wb/pkg/visitor/rectangle"
	"patterns-wb/pkg/visitor/visitor"
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
