package core

import "github.com/gin-gonic/gin"

func CreateIndex(url string) Index {
	index := Index{
		Title:       "Satheesh's API",
		Description: "This is an API for Satheesh's portfolio website and other projects.",
	}

	index.Author.Name = "Satheesh Kumar"
	index.Author.Email = "mail@satheesh.dev"
	index.Author.Websites.Portfolio = "https://www.satheesh.dev"
	index.Author.Websites.API = "https://boom.satheesh.dev"
	index.Author.Websites.Assets = "https://assets.satheesh.dev"
	index.Author.Websites.Playground = "https://play.satheesh.dev"

	index.API.Me.Info = url + "/me"
	index.API.Me.Skills = url + "/me/skills"
	index.API.Me.Services = url + "/me/services"
	index.API.Me.Projects = url + "/me/projects"
	index.API.Me.Education = url + "/me/education"
	index.API.Me.Experience = url + "/me/experience"
	index.API.Me.Testimonials = url + "/me/testimonials"

	index.API.Games.Flames.URL = url + "/games/flames"
	index.API.Games.Flames.Body = gin.H{
		"name":         "name1",
		"partner_name": "name2",
	}

	return index
}
