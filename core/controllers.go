package core

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
    scheme := "http"

    if c.Request.TLS != nil {
        scheme = "https"
    }

	c.JSON(200, gin.H{
		"title":       "Satheesh's API",
		"description": "This is an API for Satheesh's portfolio website and other projects.",
		"author": gin.H{
			"name":  "Satheesh Kumar",
			"email": "mail@satheesh.dev",
			"websites": gin.H{
				"portfolio":  "https://www.satheesh.dev",
				"api":        "https://boom.satheesh.dev",
				"assets":     "https://assets.satheesh.dev",
				"playground": "https://play.satheesh.dev",
			},
		},
		"api": gin.H{
			"me": gin.H{
				"GET_info":       scheme+"://"+c.Request.Host+"/me",
				"GET_skills":     scheme+"://"+c.Request.Host+"/me/skills",
				"GET_services":   scheme+"://"+c.Request.Host+"/me/services",
				"GET_projects":   scheme+"://"+c.Request.Host+"/me/projects",
				"GET_education":  scheme+"://"+c.Request.Host+"/me/education",
				"GET_experience": scheme+"://"+c.Request.Host+"/me/experience",
				"GET_reviews":    scheme+"://"+c.Request.Host+"/me/reviews",
			},
			"games": gin.H{
				"POST_flames": gin.H{
					"url": scheme+"://"+c.Request.Host+"/games/flames",
					"body": gin.H{
						"name":         "name1",
						"partner_name": "name2",
					},
				},
			},
		},
	})
}
