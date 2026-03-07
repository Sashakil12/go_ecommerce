package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger)
	wrappedMux := manager.WrapMux(mux)
	InitRoutes(mux, manager)
	fmt.Println("srvr on 3000")

	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {
		fmt.Println("failed to start go server!", err)
	}
}
