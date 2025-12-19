package main

import (
	"fmt"
	"my_web_app/backend/config"
	"my_web_app/backend/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup router
	router := routes.SetupRouter(cfg)

	// Start server
	fmt.Printf("Server starting on port %s\n", cfg.ServerPort)
	if err := router.Run(cfg.ServerPort); err != nil {
		panic(err)
	}
}
