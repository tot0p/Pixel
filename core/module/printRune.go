package module

import (
	"errors"
	"fmt"
	"os"
	"unicode/utf8"
)

func PrintRune(r rune) {
	l := utf8.RuneLen(r)
	if l == -1 {
		fmt.Println("Error while printing the rune: ", errors.New("invalid rune"))
		return
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	if err != nil {
		fmt.Println("Error while printing the rune: ", err)
		return
	}
}
