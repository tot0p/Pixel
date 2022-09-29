package core

import (
	"Pixel/core/module"
	"fmt"
)

type Interpreter struct {
	Code  []uint8
	stack *Stack
}

func CreateInterpreter(code []uint8) *Interpreter {
	return &Interpreter{code, &Stack{0, nil}}
}

func (i Interpreter) String() string {
	return fmt.Sprint(*i.stack)
}

func (i *Interpreter) Run() {
	for _, v := range i.Code {
		i.Call(v)
	}
	fmt.Println("\nreturn 0")
}

func (i *Interpreter) Call(v uint8) {
	if v == 0 {
		return
	}
	if v >= 1 && v <= 128 {
		i.stack.Push(v)
		return
	} else if v >= 129 && v <= 255 {
		switch v {
		case 129:
			r := int32(i.stack.Pop())
			module.PrintRune(r)
		}
	}
}
