package core

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

func LoadImage(path string) []uint8 {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error while opening the file: ", err)
		return []uint8{}
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error while decoding the file: ", err)
		return []uint8{}
	}
	bounds := img.Bounds()
	nrgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(nrgba, nrgba.Bounds(), img, bounds.Min, draw.Src)
	return nrgba.Pix
}
