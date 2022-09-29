package core

import (
	"Pixel/core/module"
	"fmt"
)

type Pixel struct {
	R, G, B, A uint8
}

func (p Pixel) String() string {
	return fmt.Sprintf("R:%d G:%d B:%d A:%d", p.R, p.G, p.B, p.A)
}

func (p *Pixel) Call(n int, stack *Stack) int {
	if p.A < 255 {
		return n + 1
	}
	switch p.R {
	case 255:
		stack.Push(p.G + p.B)
	case 128:
		module.PrintRune(int32(stack.Pop()))
	}
	return n + 1
}
