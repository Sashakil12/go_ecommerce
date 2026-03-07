package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {

	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)

}
func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	n := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		n = middlewares[i](n)
	}
	return n

}
func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	n := handler
	if len(mngr.globalMiddlewares) > 0 {
		for i := len(mngr.globalMiddlewares) - 1; i >= 0; i-- {
			n = mngr.globalMiddlewares[i](n)
		}
	}
	return n
}
