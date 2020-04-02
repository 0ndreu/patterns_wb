package main

import (
	"fmt"
	"testing"

	"github.com/0ndreu/patterns-wb/pkg/visitor/areacalculator"
	"github.com/0ndreu/patterns-wb/pkg/visitor/circle"
	"github.com/0ndreu/patterns-wb/pkg/visitor/rectangle"
	"github.com/0ndreu/patterns-wb/pkg/visitor/visitor"
)

var (
	okCircle      = 78.53981633974483
	okRectangle   = 80.0
	zeroCircle    = "there is no circle"
	zeroRectangle = "there is no rectangle"
)

func TestOkCircle(t *testing.T) {
	area := areacalculator.NewArea()
	circle, err := circle.NewCircle(area, 5)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		visCircle := visitor.NewVisitor(circle)
		a := visCircle.Accept()
		if a != okCircle {
			t.Errorf("area != circle(5)")
		}
	}
}

func TestOkRectangle(t *testing.T) {
	area := areacalculator.NewArea()
	rect, err := rectangle.NewRectangle(area, 8, 10)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		visRect := visitor.NewVisitor(rect)
		res := visRect.Accept()
		if res != okRectangle {
			t.Errorf("area != rectangle(8, 10)")
		}
	}
}

func TestZeroCircle(t *testing.T) {
	area := areacalculator.NewArea()
	_, err := circle.NewCircle(area, 0)
	if err != nil {
		return
	} else {
		t.Errorf(zeroCircle)
	}
}

func TestZeroRectangle(t *testing.T) {
	area := areacalculator.NewArea()
	_, err := rectangle.NewRectangle(area, 0, 0)
	if err != nil {
		return
	} else {
		t.Errorf(zeroRectangle)
	}
}
