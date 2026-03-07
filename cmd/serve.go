package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.CorsWithPreflight)
	InitRoutes(mux, manager)
	fmt.Println("srvr on 3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("failed to start go server!", err)
	}
}
