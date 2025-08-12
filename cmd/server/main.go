package main

import (
	"log"

	"github.com/SurekshyaS/distributed/internal/handlers"
	"github.com/SurekshyaS/distributed/internal/monitoring"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Prometheus metrics endpoint
	r.GET("/metrics", monitoring.PrometheusHandler())

	// Health and sample endpoints
	r.GET("/ping", handlers.PingHandler)
	r.GET("/healthz", handlers.PingHandler)
	// Example: Users endpoint
	r.GET("/users", handlers.GetUsersHandler)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
	
}
