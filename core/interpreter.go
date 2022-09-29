package core

import (
	"fmt"
)

type Interpreter struct {
	Code  []*Pixel
	stack *Stack
}

func CreateInterpreter(code []*Pixel) *Interpreter {
	if code == nil {
		return nil
	}
	return &Interpreter{code, &Stack{0, nil}}
}

func (i Interpreter) String() string {
	return fmt.Sprint("Stack : ", *i.stack, "\nle code : ", i.Code)
}

func (i *Interpreter) Run() {
	for k, v := range i.Code {
		v.Call(k)
	}
	fmt.Println("\nreturn 0")
}

//func (i *Interpreter) Call(p *Pixel) {
//	if v == 0 {
//		return
//	}
//	if v >= 1 && v <= 128 {
//		i.stack.Push(v)
//		return
//	} else if v >= 129 && v <= 255 {
//		switch v {
//		case 129:
//			r := int32(i.stack.Pop())
//			module.PrintRune(r)
//		}
//	}
//}
