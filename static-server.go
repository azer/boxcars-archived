package boxcars

import (
	"errors"
	"net/http"
)

type StaticServer struct {
	handler      http.Handler
	hasCustom404 bool
	custom404    string
}

func (server *StaticServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler := StaticHandler{}
	handler.request = request
	handler.ResponseWriter = writer
	handler.server = server
	server.handler.ServeHTTP(&handler, request)
}

type StaticHandler struct {
	http.ResponseWriter
	request *http.Request
	server  *StaticServer
	is404   bool
}

func (handler *StaticHandler) WriteHeader(n int) {
	if n == http.StatusNotFound && handler.server.hasCustom404 {
		debug("Serving %s as custom 404.", handler.server.custom404)
		handler.ResponseWriter.Header().Set("Content-Type", "text/html")
		http.ServeFile(handler.ResponseWriter, handler.request, handler.server.custom404)
		return
	}

	handler.ResponseWriter.WriteHeader(n)
}

func (handler *StaticHandler) Write(b []byte) (int, error) {
	if handler.is404 {
		return 0, errors.New("404 status written, will serve custom 404 page")
	}

	return handler.ResponseWriter.Write(b)
}
