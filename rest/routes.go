package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProductsHandler)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProductsHandler)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductsByIdHandler)))
}
