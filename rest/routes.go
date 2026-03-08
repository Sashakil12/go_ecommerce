package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProductsHandler)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProductsHandler), middleware.Authentication))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductsByIdHandler)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProductByIdHandler), middleware.Authentication))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProductByIdHandler), middleware.Authentication))
	//users
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUsersHandler)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.LoginHandler)))

}
