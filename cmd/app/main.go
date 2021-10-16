package main

import (
	"context"
	"log"

	apiServer "github.com/AlexanderGureev/go-service-template/internal/server"
	env "github.com/AlexanderGureev/go-service-template/pkg/env"
)

func init() {
	env.Load()
}

func main() {
	server := apiServer.Create(context.Background())

	PORT := env.Get("PORT", "3000")

	if err := server.Listen(PORT); err != nil {
		log.Fatal(err)
	}

	log.Printf("server listen on port %s", PORT)
}
