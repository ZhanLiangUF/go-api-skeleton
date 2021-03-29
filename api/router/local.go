package router //import github.com/ZhanLiangUF/go-api-skeleton/api/router

import (
	"net/http"

	"github.com/ZhanLiangUF/go-api-skeleton/api/httputils"
)

// RouteWrapper wraps a route to provide it extra functionality. This function type describes a function with the same parameters and same result type
type RouteWrapper func(r Route) Route

// localRoute implements Route
type localRoute struct {
	method  string
	path    string
	handler httputils.APIFunc
}

// Handler returns the APIFunc to be used to wrap it in middleware
func (l localRoute) Handler() httputils.APIFunc {
	return l.handler
}

// Method returns the http method that the route responds to
func (l localRoute) Method() string {
	return l.method
}

// Path returns the subpath
func (l localRoute) Path() string {
	return l.path
}

// NewRoute initializes a new local route for the router
func NewRoute(method, path string, handler httputils.APIFunc, opts ...RouteWrapper) Route {
	var r Route = localRoute{method, path, handler}
	for _, o := range opts {
		r = o(r)
	}
	return r
}

// NewGetRoute initializes a new route with the http method GET
func NewGetRoute(path string, handler httputils.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodGet, path, handler, opts...)
}
