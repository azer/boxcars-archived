package boxcars

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

var (
	config string
	doc map[string]string
	handlers map[string]http.Handler
	initialized bool = false
)

func Load (filename string) {
	debug("Loading %s", filename)

	config = filename
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		debug("Failed to read %s", filename)
		return;
	}

	err = json.Unmarshal(content, &doc)
	return
}

func Sites () (table map[string]http.Handler) {

	if initialized {
		table = handlers
		return
	}

	debug("Initializing the file servers")

	table = make(map[string]http.Handler)
	handlers = table
	initialized = true

	for site, config := range doc {

		if matches, _ := regexp.MatchString("^/", config); matches {
			debug("A file server to serve %s at %s", config, site)
			table[site] = http.FileServer(http.Dir(config))
			continue
		}

		debug("A proxy server to serve %s at %s", config, site)
		dest, _ := url.Parse(addProtocol(config))
		table[site] = httputil.NewSingleHostReverseProxy(dest)
	}

	return
}

func addProtocol (url string) string {
	if matches, _ := regexp.MatchString("^\\w+://", url); !matches {
		return fmt.Sprintf("http://%s", url)
	}

	return url
}
