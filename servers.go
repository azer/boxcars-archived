package boxcars

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

func ReverseProxyServer(uri string) http.Handler {
	debug("Returning a reverse proxy server for %s.", uri)
	dest, _ := url.Parse(addProtocol(uri))
	return httputil.NewSingleHostReverseProxy(dest)
}

func newStaticServer(uri string, hasCustom404 bool, custom404 string) http.Handler {
	debug("Returning a static server for %s", uri)
	return &StaticServer{http.FileServer(http.Dir(uri)), hasCustom404, custom404}
}

func newSingleFileServer(uri string) http.Handler {
	debug("Returning a single file server for %s", uri)
	return &SingleFileServer{uri}
}

func addProtocol(url string) string {
	if matches, _ := regexp.MatchString("^\\w+://", url); !matches {
		return fmt.Sprintf("http://%s", url)
	}

	return url
}
