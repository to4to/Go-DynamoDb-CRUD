package routes

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {

	return &Router{
		config: Newconfig().SetTimeOut(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}

}

func (r *Router) SetRouters() *chi.Mux {}

func (r *Router) setConfigsRouter() {}
