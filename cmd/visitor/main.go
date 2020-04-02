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
	circle, err := circle.NewCircle(area, 5)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		visCircle := visitor.NewVisitor(circle)
		a := visCircle.Accept()
		fmt.Println(a)
	}
	rect, err := rectangle.NewRectangle(area, 8, 10)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		visRect := visitor.NewVisitor(rect)
		res := visRect.Accept()
		fmt.Println(res)
	}

}
