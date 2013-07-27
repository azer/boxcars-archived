package boxcars

import (
	"net/http"
	"strings"
	"fmt"
)

func matchingHandlerOf (url, hostname string, handlers Handlers) (result http.Handler, found bool) {

	if handlers == nil {
		return nil, false
	}

	for pattern, handler := range handlers {
		debug("Iterating patterns: %s", pattern)

		if pattern == "*" {
			continue
		}

		if len(url) >= len(pattern) && url[0:len(pattern)] == pattern {
			debug("Matched %s%s with the handler attached to %s.", hostname, url, pattern)
			found = true
			result = http.StripPrefix(pattern, handler.server)
		}
	}

	if handler, hasDefaultHandler := handlers["*"]; !found && hasDefaultHandler {
		debug("Matched %s%s with default handler.", hostname, url)
		result = handler.server
		found = true
	}

	return result, found
}

func matchingServerOf (host, url string) (result http.Handler, found bool) {

	hostname := hostnameOf(host)
	wildcard := wildcardOf(hostname)

	result, found = matchingHandlerOf(url, hostname, sites[hostname])

	if !found {
		if _, hasWildcard := sites[wildcard]; hasWildcard {
			debug("Matching the wildcard %s", wildcard)
			result, found = matchingHandlerOf(url, hostname, sites[wildcard])
		} else {
			debug("Nothing attached to %s or %s", hostname, wildcard)
		}
	}

	debug("returning matching server")

	return result, found
}

func hostnameOf (host string) string {
	hostname := strings.Split(host, ":")[0]

	if len(hostname) > 4 && hostname[0:4] == "www." {
		hostname = hostname[4:]
	}

	return hostname
}

func wildcardOf (hostname string) string {
	parts := strings.Split(hostname, ".")

	if len(parts) < 3 {
		return fmt.Sprintf("*.%s", hostname)
	}

	parts[0] = "*"
	return strings.Join(parts, ".")

}
