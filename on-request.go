package boxcars

import (
	"net/http"
)

func OnRequest(w http.ResponseWriter, r *http.Request) {
	debug("Routing %s%s", r.Host, r.URL)

	server, found := matchingServerOf(r.Host, r.URL.String())

	if found {
		server.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}
