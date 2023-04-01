package main

import (
    "log"
    "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    router := gin.Default()
    return router
}

func main() {
    router := setupRouter()
    log.Fatal(router.Run(":8080"))
}
