package core

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
)

type imgInCompile struct {
	Img *image.RGBA
	con []byte
}

func Compile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error while opening the file: ", err)
		return
	}
	defer file.Close()
	con, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error while reading the file: ", err)
		return
	}
	img := imgInCompile{con: con}
	img.Compile()
	savePng(path, img.Img)
}

func (i *imgInCompile) Compile() {
	matrix := [][]string{}
	i.con = []byte(strings.ReplaceAll(string(i.con), "\r", ""))
	temp := strings.Split(string(i.con), "\n")
	for _, v := range temp {
		if v != "" {
			matrix = append(matrix, strings.Split(v, ","))
			if len(matrix[len(matrix)-1]) != len(matrix[0]) {
				fmt.Println("The image is not a rectangle")
				return
			} else if (len(matrix[len(matrix)-1]) % 4) != 0 {
				fmt.Println("Wrong number of nb")
				return
			}
		}
	}
	i.Img = image.NewRGBA(image.Rect(0, 0, len(matrix[0])/4, len(matrix)))
	pix := []uint8{}
	for _, v := range matrix {
		for _, v2 := range v {
			t, err := strconv.Atoi(v2)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if t > 255 || t < 0 {
				fmt.Println("The number is not between 0 and 255")
				return
			}
			pix = append(pix, uint8(t))
		}
	}
	i.Img.Pix = pix
}

func savePng(path string, img image.Image) {
	f, err := os.Create(path + ".png")
	if err != nil {
		fmt.Println("Error while creating the file: ", err)
		return
	}
	png.Encode(f, img)
	f.Close()
}
