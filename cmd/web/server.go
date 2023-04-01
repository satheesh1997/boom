package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/satheesh1997/boom/games"
)

func setupRouter() *gin.Engine {
    router := gin.Default()

    // Controllers
    gamesController := games.NewController()

    // Routes for /games
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
