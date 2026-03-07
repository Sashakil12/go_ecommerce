package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.CorsWithPreflight, middleware.Logger)
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProductsHandler)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProductsHandler)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductsByIdHandler)))
	fmt.Println("srvr on 3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("failed to start go server!", err)
	}
}
