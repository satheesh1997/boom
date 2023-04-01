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

// @Summary Compute Flames
// @Description Compute Flames between two name using FLAMES game rule
// @Tags games
// @Accept  json
// @Produce  json
// @Param body body ComputeFlamesRequest true "Request Body"
// @Success 200 {object} ComputeFlamesResponse
// @Failure 400 {object} ErrorResponse
// @Router /games/flames [post]
func (controller *Controller) ComputeFlames(c *gin.Context) {
    var req ComputeFlamesRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := gameService.ComputeFlames(req.Name, req.PartnerName)
    c.JSON(http.StatusOK, gin.H{"result": result})
}
