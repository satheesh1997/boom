package me

import (
	"github.com/gin-gonic/gin"
)

func (controller Controller) GetInfo(c *gin.Context) {
	c.JSON(200, CreateInfo())
}

func (controller Controller) GetSkills(c *gin.Context) {
	c.JSON(200, CreateSkills())
}

func (controller Controller) GetServices(c *gin.Context) {
	c.JSON(200, CreateMyServices())
}

func (controller Controller) GetProjects(c *gin.Context) {
	c.JSON(200, CreateProjects())
}

func (controller Controller) GetEducation(c *gin.Context) {
	c.JSON(200, CreateEducation())
}

// Note: This will be replaced with linkedin api to fetch the data
// from next version
func (controller Controller) GetExperience(c *gin.Context) {
	c.JSON(200, []gin.H{
		gin.H{
			"id":              1,
			"title":           "Software Engineer",
			"type":            "Internship",
			"company":         "HackerEarth",
			"companyLocation": "Bangalore, India",
			"date": gin.H{
				"start":    "May 2018",
				"end":      "Feb 2019",
				"duration": "10 months",
			},
			"description": "Worked as a software engineer intern in the recruit team.",
			"contributions": []string{
				"Framework for the setter app that reduced the time taken for introducing a new problem by ***75%***.",
				"Optimised page load times to less than 2 secs.",
				"Introduced Project problem type in content creation interface.",
				"Developed internal tools that help the on-call & support team resolve the on-call issues without much effort.",
				"Coordinated with the Support team to resolve customer issues.",
			},
			"skills": []string{
				"HTML",
				"RabbitMQ",
				"MySQL",
				"MongoDB",
				"Redis",
				"InfluxDB",
				"Django",
				"Python",
				"Jquery",
			},
		},
		gin.H{
			"id":              2,
			"title":           "Software Engineer",
			"type":            "Full Time",
			"company":         "HackerEarth",
			"companyLocation": "Bangalore, India",
			"date": gin.H{
				"start":    "Apr 2019",
				"end":      "Feb 2021",
				"duration": "1 year 11 months",
			},
			"description": "Worked as a Backend engineer in core applications & recruit team.",
			"contributions": []string{
				"Coordinated with the UI devs & created REST API for the latest UI components.",
				"Helped UI devs by getting hands on with them.",
				"Migrated from pusher to firebase for realtime notifications.",
				"Introduced FullstackProblem framework that helps customers evaluate full-stack (Django, React, ..etc) skills.",
				"Reduced slow page(s) load time by ***90%***.",
				"Made integration tests run in parallel, this reduced the time developers need to wait for PR to get merged.",
				"Worked on a project evaluation engine that helps evaluate the python & csharp skills of a candidate.",
				"Introduced new IDE for candidates to work on project question types.",
				"Mentored & onboarded the new Application engineering team.",
			},
			"skills": []string{
				"Firebase",
				"Django Rest Framework",
				"React",
				"Datadog",
				"Sentry",
			},
		},
		gin.H{
			"id":              3,
			"title":           "Senior Software Engineer",
			"type":            "Full Time",
			"company":         "HackerEarth",
			"companyLocation": "Bangalore, India",
			"date": gin.H{
				"start":    "Feb 2021",
				"end":      "Apr 2022",
				"duration": "1 year 3 months",
			},
			"description": "Worked as Senior backend engineer in the core applications team.",
			"contributions": []string{
				"Migrated from Fargate to EC2 to improve the performance in IDE services.",
				"Played a major role in the process of making the HackerEarth codebase python3 compatible.",
				"Reduced cost incurred from firebase by ***90%***.",
				"Created infra on AWS for a service that detects code plagiarism from the internet.",
				"Introduced DevOpsProblem framework that helps customers evaluate DevOps (Docker, k8s, ..etc) skills.",
				"Worked along with platform engineers on architectural design for the new analytics flow.",
				"Initiative to write backend test cases. Added 100% coverage for an django app.",
				"Coordinated with the Application engineering team to make them resolve customer issues on time.",
			},
			"skills": []string{
				"Golang",
				"Nose",
				"AWS EC2",
				"AWS Fargate",
				"SSH",
				"Docker",
				"AWS SSM",
				"Django",
			},
		},
		gin.H{
			"id":              4,
			"title":           "Staff Software Engineer",
			"type":            "Full Time",
			"company":         "HackerEarth",
			"companyLocation": "Bangalore, India",
			"date": gin.H{
				"start":    "Apr 2022",
				"end":      "Present",
				"duration": "1 year 3 months",
			},
			"description": "Working as an IC in platform team.",
			"contributions": []string{
				"Created a python client to communicate with analytics service using TDD methodology, this stand the project that has highest coverage of ***81%***.",
				"Implemented an end to end UI using react for orchestration service.",
				"Developed and maintaining the analytics service.",
				"Reduced the monthly cost spent on ECR to **50%**.",
			},
			"skills": []string{
				"Mock",
				"React",
				"Puppeteer",
				"AWS DynamoDB",
				"Semantic UI",
				"AWS ECR",
			},
		},
	})
}

// Note: This will be replaced with linkedin api to fetch the data
// from next version
func (controller Controller) GetTestimonials(c *gin.Context) {
	c.JSON(200, CreateTestimonials())
}
