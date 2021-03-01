package api // import "github.com/ZhanLiangUF/go-flights/api"

import "github.com/ZhanLiangUF/go-flights/api/httputils"

type Router interface {
	Routes() []Route
}

type Route interface {
	// Handler returns the raw function to create the http handler
	Handler() httputils.APIFunc
	Method() string
	Path() string
}