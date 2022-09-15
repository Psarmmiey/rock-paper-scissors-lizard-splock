package routes

import (
	"rock-paper/internal/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/choices", controllers.GetChoices)
	g.GET("/choice", controllers.GetSingleChoice)
	g.POST("/play", controllers.Play)
	g.POST("/multiplayer", controllers.Multiplayer)
	g.GET("/reset", controllers.ResetGameManually)
}
