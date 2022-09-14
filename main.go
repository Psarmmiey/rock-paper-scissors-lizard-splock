package main

import (
	"rock-paper/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	public := router.Group("/")
	routes.PublicRoutes(public)
	router.Run()
}
