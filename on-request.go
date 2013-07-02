package boxcars

import (
	"fmt"
	"net/http"
	"strings"
)

func OnRequest(w http.ResponseWriter, r *http.Request) {
	hostname := strings.Split(r.Host, ":")[0]
	table := Sites()
	server, defined := table[hostname]

	debug("New request to %s/%s", hostname, r.URL)

	if defined {
		server.ServeHTTP(w, r)
	}

	fmt.Fprintf(w, "404 - Not found.")
}
