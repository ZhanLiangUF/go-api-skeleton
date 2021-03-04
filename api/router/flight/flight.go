package flight // import "github.com/ZhanLiangUF/go-flights/api/router/flights"
import (
	router "github.com/ZhanLiangUF/go-flights/api/router"
)

type flightRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	r := &flightRouter{}
	r.initRoutes()
	return r
}

func (r *flightRouter) Routes() []router.Route {
	return r.routes
}

func (r *flightRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/flights/get", r.getFlightsGet),
	}
}
