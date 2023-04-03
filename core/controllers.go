package core

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
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
				"GET_info":       "http://localhost:8080/me",
				"GET_skills":     "http://localhost:8080/me/skills",
				"GET_services":   "http://localhost:8080/me/services",
				"GET_projects":   "http://localhost:8080/me/projects",
				"GET_education":  "http://localhost:8080/me/education",
				"GET_experience": "http://localhost:8080/me/experience",
				"GET_reviews":    "http://localhost:8080/me/reviews",
			},
			"games": gin.H{
				"POST_flames": gin.H{
					"url": "http://localhost:8080/games/flames",
					"body": gin.H{
						"name":         "name1",
						"partner_name": "name2",
					},
				},
			},
		},
	})
}
