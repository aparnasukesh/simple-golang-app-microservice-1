package main

import (
	"log"
	"microservice-one/config"
	"microservice-one/internals/di"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error loading config file")
	}

	userHandler, err := di.InitResources(cfg)
	if err != nil {
		log.Fatalf("Error happened while resource initialization: %v", err)
	}
	r := gin.Default()
	user := r.Group("/user")
	userHandler.Routes(user)
	r.Run(":8080")
}
