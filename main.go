package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/satheesh1997/boom/core"
	"github.com/satheesh1997/boom/games"
	"github.com/satheesh1997/boom/me"
)

var (
	ginLambda *ginadapter.GinLambda
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func inRelease() bool {
	if release := os.Getenv("GIN_MODE"); release == "release" {
		return true
	}
	return false
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://www.satheesh.dev", "https://play.satheesh.dev", "http://localhost:3002"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	// Controllers
	gamesController := games.NewController()
	meController := me.NewController()

	// Route for /
	router.GET("/", core.IndexHandler)

	// Route for /health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Route for /me
	meGroup := router.Group("/me")
	{
		meGroup.GET("", meController.GetInfo)
		meGroup.GET("/skills", meController.GetSkills)
		meGroup.GET("/services", meController.GetServices)
		meGroup.GET("/projects", meController.GetProjects)
		meGroup.GET("/education", meController.GetEducation)
		meGroup.GET("/experience", meController.GetExperience)
		meGroup.GET("/testimonials", meController.GetTestimonials)
	}

	// Route for /games/*
	gamesGroup := router.Group("/games")
	{
		gamesGroup.POST("/flames", gamesController.ComputeFlames)
	}

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    "API_NOT_FOUND",
			"message": "Please reach out to the developer at mail@satheesh.dev for documentation.",
		})
	})

	return router
}

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func init() {
	if inLambda() || inRelease() {
		ginLambda = ginadapter.New(setupRouter())
	}
}

func main() {
	router := setupRouter()

	if inLambda() || inRelease() {
		lambda.Start(lambdaHandler)
		return
	}

	log.Fatal(router.Run(":8080"))
}
