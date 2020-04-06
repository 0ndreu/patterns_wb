package reciever

import "fmt"

type Calc interface {
	Multiply()
	Sum()
}

type Calculator struct {
	X int
	Y int
}

func (c *Calculator) Multiply() {
	total := c.X * c.Y
	fmt.Println("> Multiply X & Y")
	fmt.Printf("Input: %d, %d\n", c.X, c.Y)
	fmt.Println("Output:", total)
}

func (c *Calculator) Sum() {
	total := c.X + c.Y
	fmt.Println("> Sum X & Y")
	fmt.Printf("Input: %d, %d\n", c.X, c.Y)
	fmt.Println("Output:", total)
}

func NewCalculator(x, y int) (out *Calculator) {
	out = &Calculator{
		X: x,
		Y: y,
	}
	return
}
