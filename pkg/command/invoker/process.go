package invoker

type Command interface {
	Execute()
}

type Process struct {
	Command Command
}

func (p *Process) process() {
	p.Command.Execute()
}
