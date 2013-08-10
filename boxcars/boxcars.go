package main

import (
	"flag"
	"fmt"
	"github.com/azer/boxcars"
	"os"
)

var (
	filename string
	port     int
	user_id  int
	group_id int
	secure   bool
)

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen")
	flag.BoolVar(&secure, "secure", false, "Enables secure mode to avoid running as sudo.")
	flag.IntVar(&user_id, "uid", 1000, "User id that'll own the system process.")
	flag.IntVar(&group_id, "gid", 1000, "Group id that'll own the system process.")
	flag.Parse()

	filename = flag.Arg(0)

	if filename == "" {
		fmt.Printf("Usage: boxcars config.json\n")
		os.Exit(1)
	}

	boxcars.SetFilename(filename)
	go boxcars.ReadConfig()
	go boxcars.AutoReload()

	if secure {
		go boxcars.Secure(user_id, group_id)
	}

	boxcars.Listen(port)
}
