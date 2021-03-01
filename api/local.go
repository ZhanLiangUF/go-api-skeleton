package api //import github.com/ZhanLiangUF/go-flights/api

import (
	"net/http"

	"github.com/ZhanLiangUF/go-flights/api/httputils"
)

// RouteWrapper wraps a route to provide it extra functionality. This function type describes a function with the same parameters and same result type
type RouteWrapper func(r Route) Route

// localRoute implements Route
type localRoute struct {
	method  string
	path    string
	handler httputils.APIFunc
}

func (l localRoute) Handler() httputils.APIFunc {
	return l.handler
}

func (l localRoute) Method() string {
	return l.method
}

func (l localRoute) Path() string {
	return l.path
}

// initializes a new localroute for router
func NewRoute(method, path string, handler httputils.APIFunc, opts ...RouteWrapper) Route {
	var r Route = localRoute{method, path, handler}
	for _, o := range opts {
		r = o(r)
	}
	return r
}

func NewGetRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodGet, path, handler, opts...)
}
