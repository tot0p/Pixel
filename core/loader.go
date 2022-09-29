package core

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

func LoadImage(path string) []*Pixel {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error while opening the file: ", err)
		return nil
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error while decoding the file: ", err)
		return nil
	}
	bounds := img.Bounds()
	nrgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(nrgba, nrgba.Bounds(), img, bounds.Min, draw.Src)
	return translate(nrgba.Pix)
}

func translate(pix []uint8) []*Pixel {
	var result []*Pixel
	for i := 0; i < len(pix); i += 4 {
		result = append(result, &Pixel{pix[i], pix[i+1], pix[i+2], pix[i+3]})
	}
	return result
}
