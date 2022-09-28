package core

type Interpreter struct {
	Code  []uint8
	stack *Stack
}

func CreateInterpreter(code []uint8) *Interpreter {
	return &Interpreter{code, &Stack{0, nil}}
}

func (i *Interpreter) Run() {
	for _, v := range i.Code {
		i.Call(v)
	}
}

func (i *Interpreter) Call(v uint8) {
	if v <= 0 && v <= 127 {
		i.stack.Push(v)
		return
	}
}
