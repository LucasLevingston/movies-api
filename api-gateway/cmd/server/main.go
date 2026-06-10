// @title           Movies API
// @version         1.0
// @description     REST API Gateway for the Movies microservice. Communicates with the movies-service via gRPC.
// @host            localhost:8080
// @BasePath        /
// @schemes         http

package main

import (
	"log"
	"os"

	_ "movies-api/api-gateway/docs"

	grpcclient "movies-api/api-gateway/internal/adapters/grpc"
	httphandler "movies-api/api-gateway/internal/adapters/http"
)

func main() {
	moviesAddr := getEnv("MOVIES_SERVICE_ADDR", "localhost:50051")
	port := getEnv("HTTP_PORT", "8080")

	client, err := grpcclient.NewMovieGRPCClient(moviesAddr)
	if err != nil {
		log.Fatalf("connect to movies-service: %v", err)
	}

	router := httphandler.NewRouter(client)

	log.Printf("API Gateway listening on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
