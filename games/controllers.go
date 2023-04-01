package games

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type (
    ComputeFlamesRequest struct {
        Name string `json:"name" binding:"required"`
        PartnerName string `json:"partner_name" binding:"required"`
    }
)

var (
    gameService = NewService()
)

func (controller *Controller) ComputeFlames(c *gin.Context) {
    var req ComputeFlamesRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := gameService.ComputeFlames(req.Name, req.PartnerName)
    c.JSON(http.StatusOK, gin.H{"result": result})
}
