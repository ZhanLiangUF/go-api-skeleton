package router // import "github.com/ZhanLiangUF/go-api-skeleton/api/router"

import "github.com/ZhanLiangUF/go-flights/api/httputils"

// Router defines an interface to specify the list of routes to add to the server
type Router interface {
	Routes() []Route
}

// Route defines an individual API
type Route interface {
	// Handler returns the raw function to create the http handler
	Handler() httputils.APIFunc
	Method() string
	Path() string
}
