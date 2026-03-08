package review

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Register review routes
	mux.Handle("GET /reviews", manager.With(http.HandlerFunc(h.GetReviewsHandler)))

}
