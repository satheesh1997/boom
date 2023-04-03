package main

import (
    "log"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/satheesh1997/boom/games"
    "github.com/satheesh1997/boom/me"
)

func setupRouter() *gin.Engine {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*.satheesh.dev", "http://localhost:3000"},
        AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
    }))

    // Load templates
    router.LoadHTMLGlob("templates/*")

    // Controllers
    gamesController := games.NewController()
    meController := me.NewController()

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

    // Route for /me
    meGroup := router.Group("/me")
    {
        meGroup.GET("/", meController.GetInfo)
        meGroup.GET("/skills", meController.GetSkills)
        meGroup.GET("/services", meController.GetServices)
        meGroup.GET("/projects", meController.GetProjects)
        meGroup.GET("/education", meController.GetEducation)
        meGroup.GET("/experience", meController.GetExperience)
        meGroup.GET("/reviews", meController.GetReviews)
    }

    // Route for /games/*
    gamesGroup := router.Group("/games")
    {
        gamesGroup.POST("/flames", gamesController.ComputeFlames)
    }

    // Handle 404
    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{
            "code": "API_NOT_FOUND",
            "message": "Please reach out to the developer at mail@satheesh.dev for documentation.",
        })
    })

    return router
}

func main() {
    router := setupRouter()
    log.Fatal(router.Run(":8080"))
}
