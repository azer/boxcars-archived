package boxcars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	doc map[string]string
	handlers map[string]http.Handler
	initialized bool = false
)

func Load (filename string) {
	log.Printf("Loading %s", filename)

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Failed to read %s", filename)
		os.Exit(1)
	}

	err = json.Unmarshal(content, &doc)
	return
}

func Sites () (table map[string]http.Handler) {

	if initialized {
		log.Println("Returning from cache.")
		table = handlers
		return
	}

	log.Println("Initializing handlers")

	table = make(map[string]http.Handler)
	handlers = table
	initialized = true

	for site, config := range doc {
		table[site] = http.FileServer(http.Dir(config))
	}

	return
}
