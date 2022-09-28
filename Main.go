package main

import (
	"Pixel/core"
	"flag"
	"fmt"
)

var (
	path = ""
)

func init() {
	flag.StringVar(&path, "path", "", "path to the file")
	flag.Parse()
}

func main() {
	if path == "" {
		fmt.Println("Please provide a path to the file")
		return
	}
	fmt.Println("Path to the file is: ", path)
	fmt.Println(core.LoadImage(path))
}
