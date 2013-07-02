package boxcars

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	doc map[string]string
	handlers map[string]http.Handler
	initialized bool = false
)

func Load (filename string) {
	debug("Loading %s", filename)

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		debug("Failed to read %s", filename)
		os.Exit(1)
	}

	err = json.Unmarshal(content, &doc)
	return
}

func Sites () (table map[string]http.Handler) {

	if initialized {
		debug("Returning from cache.")
		table = handlers
		return
	}

	debug("Initializing handlers")

	table = make(map[string]http.Handler)
	handlers = table
	initialized = true

	for site, config := range doc {
		table[site] = http.FileServer(http.Dir(config))
	}

	return
}
