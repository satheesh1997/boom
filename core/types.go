package core

import "github.com/gin-gonic/gin"

type (
	PostAPI struct {
		URL  string `json:"url"`
		Body gin.H  `json:"body"`
	}

	Index struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Author      struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Websites struct {
				Portfolio  string `json:"portfolio"`
				API        string `json:"api"`
				Assets     string `json:"assets"`
				Playground string `json:"playground"`
			} `json:"websites"`
		} `json:"author"`
		API struct {
			Me struct {
				Info         string `json:"GET_info"`
				Skills       string `json:"GET_skills"`
				Services     string `json:"GET_services"`
				Projects     string `json:"GET_projects"`
				Education    string `json:"GET_education"`
				Experience   string `json:"GET_experience"`
				Testimonials string `json:"GET_testimonials"`
			} `json:"me"`
			Games struct {
				Flames PostAPI `json:"POST_flames"`
			} `json:"games"`
		} `json:"api"`
	}
)
