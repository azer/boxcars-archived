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
		return
	}

	var parsed map[string]string

	err = json.Unmarshal(content, &parsed)

	if err != nil {
		debug("Failed to parse %s", filename)
		return
	}

	doc = parsed
	initialized = false

	return
}

func Sites () (table map[string]http.Handler) {

	if initialized {
		table = handlers
		return
	}

	table = make(map[string]http.Handler)
	handlers = table
	initialized = true

	for site, config := range doc {
		if isLocalPath(config) {
			debug("%s is set to serve %s", site, config)
			table[site] = http.FileServer(http.Dir(config))
			continue
		}

		debug("%s is set to reverse proxy to %s", site, config)
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

func isLocalPath (config string) bool {
	matches, _ := regexp.MatchString("^/", config)
	return matches
}
