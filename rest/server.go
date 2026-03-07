package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

func Start(config *config.Config) {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger)
	wrappedMux := manager.WrapMux(mux)
	InitRoutes(mux, manager)
	fmt.Printf("srvr on %d\n", config.HttpPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.HttpPort), wrappedMux)
	if err != nil {
		fmt.Println("failed to start go server!", err)
		os.Exit(1)
	}
}
