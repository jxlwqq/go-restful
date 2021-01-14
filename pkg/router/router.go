package router

import "github.com/gorilla/mux"

type Router struct {
	*mux.Router
}

func New() *Router {
	mu := mux.NewRouter().StrictSlash(true)
	return &Router{mu}
}
