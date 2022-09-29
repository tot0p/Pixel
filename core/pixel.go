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
		if p.B == 0 {
			stack.Push(p.G)
		} else if p.G == 0 {
			stack.Push(255 * p.B)
		} else {
			stack.Push(p.G * p.B)
		}
	case 128:
		if p.G == 0 && p.B == 0 {
			module.PrintRune(int32(stack.Pop()))
		} else if p.G == 0 {
			module.PrintRune(int32(p.B * 255))
		} else if p.B == 0 {
			module.PrintRune(int32(p.G))
		} else {
			module.PrintRune(int32(p.G * p.B))
		}
	case 10:
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a + b)
	case 20:
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a - b)
	case 30:
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a * b)
	case 40:
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a / b)
	}
	return n + 1
}
