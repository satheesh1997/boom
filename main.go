package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/satheesh1997/boom/games"
)

func setupRouter() *gin.Engine {
    router := gin.Default()

    // Load templates
    router.LoadHTMLGlob("templates/*")

    // Controllers
    gamesController := games.NewController()

    // Route for /
    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })

    // Route for /health
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "OK",
        })
    })

    // Route for /games/*
    gamesGroup := router.Group("/games")
    {
        gamesGroup.POST("/flames", gamesController.ComputeFlames)
    }

    return router
}

func main() {
    router := setupRouter()
    log.Fatal(router.Run(":8080"))
}
