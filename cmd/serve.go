package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	config := config.GetConfig()
	middlewares := middleware.NewMiddlewares(config)
	//db
	dbCon, err := db.NewConnection(config.DBConfig)

	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
		return
	}
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		os.Exit(1)
		return
	}

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
	//handlers
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(userRepo, config)
	server := rest.NewServer(config,
		productHandler,
		userHandler,
	)
	server.Start()
}
