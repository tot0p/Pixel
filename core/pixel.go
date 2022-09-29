package core

import "fmt"

type Pixel struct {
	R, G, B, A uint8
}

func (p Pixel) String() string {
	return fmt.Sprintf("R:%d G:%d B:%d A:%d", p.R, p.G, p.B, p.A)
}

func (p *Pixel) Call(n int) int {
	return n + 1
}
