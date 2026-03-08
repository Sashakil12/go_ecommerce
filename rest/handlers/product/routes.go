package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProductsHandler)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProductsHandler), h.middlewares.Authentication))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProductsByIdHandler)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProductByIdHandler), h.middlewares.Authentication))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProductByIdHandler), h.middlewares.Authentication))

}
