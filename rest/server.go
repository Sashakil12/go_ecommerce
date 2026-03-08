package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	configuration  *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NewServer(configuration *config.Config, productHandler *product.Handler, userHandler *user.Handler, reviewHandler *review.Handler) *Server {
	return &Server{
		configuration:  configuration,
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
}
func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger)
	wrappedMux := manager.WrapMux(mux)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewHandler.RegisterRoutes(mux, manager)
	fmt.Printf("srvr on %d\n", server.configuration.HttpPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", server.configuration.HttpPort), wrappedMux)
	if err != nil {
		fmt.Println("failed to start go server!", err)
		os.Exit(1)
	}
}
