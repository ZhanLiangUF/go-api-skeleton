package test // import "github.com/ZhanLiangUFgo-api-skeleton/api/router/test"
import (
	router "github.com/ZhanLiangUF/go-api-skeleton/api/router"
)

type testRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	r := &testRouter{}
	r.initRoutes()
	return r
}

func (r *testRouter) Routes() []router.Route {
	return r.routes
}

func (r *testRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/flights/get", r.getFlightsGet),
	}
}
