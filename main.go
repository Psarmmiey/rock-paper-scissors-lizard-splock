package main

import (
	"rock-paper/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	public := router.Group("/")
	routes.PublicRoutes(public)
	router.Run()
}
