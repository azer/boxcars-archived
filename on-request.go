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

	debug("Routing to %s%s", r.Host, r.URL)

	if defined {
		server.ServeHTTP(w, r)
		return
	}

	fmt.Fprintf(w, "404 - Not found.")
}
