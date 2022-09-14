package main

import (
	"rock-paper/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	public := router.Group("/")
	routes.PublicRoutes(public)
	router.Run()
}
