package router

import (
	"github.com/julienschmidt/httprouter"
)

// Router is a wrapper of the httprouter.Router type
type Router struct {
	r *httprouter.Router
}

// NewGroup adds a zero overhead group of routes that share a common root path.
func (r *Router) NewGroup(path string) *RouteGroup {
	return newRouteGroup(r.r, path)
}

// Router returns the httprouter.Router of our wrapper type
func (r *Router) Router() *httprouter.Router {
	return r.r
}

// New creates a router wrapper using the httprouter.New method.
func New() *Router {
	return &Router{
		r: httprouter.New(),
	}
}
