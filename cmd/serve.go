package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	config := config.GetConfig()
	middlewares := middleware.NewMiddlewares(config)
	//repos
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	//handlers
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(userRepo, config)
	server := rest.NewServer(config,
		productHandler,
		userHandler,
	)
	server.Start()
}
