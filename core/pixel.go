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
	case 255: // add to stack
		if p.B == 0 {
			stack.Push(p.G)
		} else if p.G == 0 {
			stack.Push(255 * p.B)
		} else {
			stack.Push(p.G * p.B)
		}
	case 128: // print
		if p.G == 0 && p.B == 0 {
			module.PrintRune(int32(stack.Pop()))
		} else if p.G == 0 {
			module.PrintRune(int32(p.B * 255))
		} else if p.B == 0 {
			module.PrintRune(int32(p.G))
		} else {
			module.PrintRune(int32(int(p.G) * int(p.B)))
		}
	case 64: // reverse stack
		Reverse(stack)
	case 10: // add
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a + b)
	case 20: // sub
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a - b)
	case 30: // mul
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a * b)
	case 40: // div
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if p.B == 0 {
			b = stack.Pop()
		}
		stack.Push(a / b)
	case 50: // jump
		a, b := p.G, p.B
		return int(a) * int(b)
	case 60: // if = 0
		a, b := p.G, p.B
		if p.G == 0 {
			a = stack.Pop()
		}
		if a == 0 {
			return int(b)
		}
	}
	return n + 1
}
