package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/azer/boxcars"
)

var (
	filename string
	port int
)

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen")
	flag.Parse()

	filename = flag.Arg(0)

	if filename == "" {
		fmt.Printf("Usage: boxcars config.json\n")
		os.Exit(1)
	}

	boxcars.SetFilename(filename)
	go boxcars.ReadConfig()
	go boxcars.AutoReload()
	boxcars.Listen(port)
}
