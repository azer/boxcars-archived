package boxcars

import (
	"regexp"
	"fmt"
	"os"
)

type Handlers map[string]*Handler

func handlerOf (uri string, hasCustom404 bool, custom404 string) *Handler {
	debug("Setting up the HTTP handler that will serve %s", uri)

	handler := &Handler{ false, false, uri, nil }
	isStatic := isLocalPath(uri)

	if isStatic && isSingleFile(uri) {
		handler.isStatic = true
		handler.server = newSingleFileServer(uri)
  } else if isStatic {
		handler.isStatic = true
		handler.server = newStaticServer(uri, hasCustom404, custom404)
	} else {
		handler.isReverseProxy = true
		handler.server = ReverseProxyServer(uri)
	}

	return handler
}

func handlersOf (options map[string]string) Handlers {
	handlers := make(Handlers)

	custom404, hasCustom404 := options["*"]

	if hasCustom404 {
		hasCustom404 = isLocalPath(custom404)
	}

	if hasCustom404 {
		custom404 = fmt.Sprintf("%s/index.html", custom404)
	}

	for path, uri := range options {
		handlers[path] = handlerOf(uri, hasCustom404, custom404)
	}

	return handlers
}

func isLocalPath (config string) bool {
	matches, _ := regexp.MatchString("^/", config)
	return matches
}

func isSingleFile (uri string) bool {
	f, err := os.Open(uri)

	if err != nil {
		return false
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return false
	case mode.IsRegular():
		return true
	}

	return false
}
