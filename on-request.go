package boxcars

import (
	"fmt"
	"net/http"
	"strings"
)

func OnRequest(w http.ResponseWriter, r *http.Request) {
	debug("Routing %s%s", r.Host, r.URL)

	server, found := handlerOf(r)

	if found {
		server.ServeHTTP(w, r)
		return
	}

	fmt.Fprintf(w, "404 - Not found.")
}

func handlerOf (request *http.Request) (http.Handler, bool) {
	table := Sites()

	hostname := hostnameOf(request)
	handler, defined := table[hostname]

	if defined {
		return handler, true
	}

	return nil, false
}

func hostnameOf (request *http.Request) string {
	hostname := strings.Split(request.Host, ":")[0]

	if hostname[0:4] == "www." {
		hostname = hostname[4:]
	}

	return hostname
}
