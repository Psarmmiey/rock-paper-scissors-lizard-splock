package main

import (
	logger "rock-paper/internal/log"
	"rock-paper/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	var Secret = []byte("secret")
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(sessions.Sessions("session", cookie.NewStore(Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)
	router.Run(":4000")

	logger.LogDebug("-----Application started-----")
}
