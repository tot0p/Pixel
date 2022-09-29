package main

import (
	"Pixel/core"
	"flag"
	"fmt"
)

var (
	path    = ""
	compile = false
)

func init() {
	flag.BoolVar(&compile, "compile", false, "Compile the file")
	flag.StringVar(&path, "path", "", "path to the file")
	flag.Parse()
}

func main() {
	if path == "" {
		fmt.Println("Please provide a path to the file")
		return
	}
	if compile {
		core.Compile(path)
		return
	} else {
		fmt.Println("Path to the file is: ", path)
		i := core.CreateInterpreter(core.LoadImage(path))
		i.Run()
		fmt.Println(i)
	}
}
