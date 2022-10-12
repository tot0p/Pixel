package main

import (
	"Pixel/core"
	"flag"
	"fmt"
)

var (
	path    = ""
	compile = flag.Bool("c", false, "Compile the file")
	debug   = flag.Bool("d", false, "Debug the file")
)

func init() {
	flag.Parse()
}

func main() {
	path = flag.Arg(0)
	if path == "" {
		fmt.Println("Please provide a path to the file")
		return
	}
	if *compile {
		core.Compile(path)
		return
	} else {
		fmt.Println("Path to the file is: ", path)
		i := core.CreateInterpreter(core.LoadImage(path))
		i.Run()
		if *debug {
			fmt.Println(i)
		}
	}
}
