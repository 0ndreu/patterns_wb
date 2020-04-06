package sum

type calc interface {
	Sum()
}

type Command interface {
	Execute()
}

type SumCommand struct {
	Calc calc
}

func (c *SumCommand) Execute() {
	c.Calc.Sum()
}
