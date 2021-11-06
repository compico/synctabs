package hs

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	tpath = "templates/"
)

type HttpServer struct {
	Server *http.Server
	Router *httprouter.Router
	HandlerFuncs *[]httprouter.Handle
}

const (
	GET = "GET"
	POST = "POST"
	HEAD = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH = "PATCH"
	PUT = "PUT"
	DELETE = "DELETE"
)

func (s *HttpServer) AddHandle(method string, path string, handle httprouter.Handle)  {
	switch method { 
		case GET:
			s.Router.GET(path,handle)
		case POST:
			s.Router.POST(path,handle)
		case HEAD:
			s.Router.HEAD(path,handle)
		case OPTIONS:
			s.Router.OPTIONS(path,handle)
		case PATCH:
			s.Router.PATCH(path,handle)
		case PUT:
			s.Router.PUT(path,handle)
		case DELETE:
			s.Router.DELETE(path,handle)
	}
}