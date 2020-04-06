package multiply

type calc interface {
	Multiply()
}

type Command interface {
	Execute()
}

type MultiplyCommand struct {
	Calc calc
}

func (c *MultiplyCommand) Execute() {
	c.Calc.Multiply()
}
