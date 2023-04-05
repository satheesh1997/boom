package me

import (
	"github.com/google/uuid"
)

func CreateInfo() *Info {
	info := &Info{
		FirstName:      "Satheesh",
		LastName:       "Kumar",
		FullName:       "Satheesh Kumar",
		Bio:            "Engineer @HackerEarth ||  Open Source Contributor  || Designer  ||  Developer || Mentor",
		DateOfBirth:    "10-10-1997",
		Address:        "Bengaluru, Karnataka, India",
		ImageURL:       "https://assets.satheesh.dev/images/about-image.jpg",
		ThumbnailURL:   "https://assets.satheesh.dev/images/formal-image.jpg",
		Languages:      []string{"English", "Tamil", "Hindi"},
		PhoneNumbers:   []string{"+91-9597264229"},
		EmailAddresses: []string{"mail@satheesh.dev"},
	}
	info.SocialProfile.Whatsapp = "https://wa.me/919597264229"
	info.SocialProfile.LinkedIn = "https://www.linkedin.com/in/satheesh1997/"
	info.SocialProfile.Twitter = "https://twitter.com/git_push_tweet"
	info.SocialProfile.Github = "https://github.com/satheesh1997"
	info.SocialProfile.Medium = "https://esc-wq.medium.com/"
	return info
}

func CreateEducation() []Education {
	return []Education{
		{
			ID:          uuid.NewString(),
			Institution: "Amrita Vidyalayam",
			Degree:      "HSC",
			Details:     "Maths, Physics, Chemistry, Computer Science",
			Year:        "2013 - 2015",
		},
		{
			ID:          uuid.NewString(),
			Institution: "Sri Krishna College of Engineering and Technology",
			Degree:      "B.E",
			Details:     "Computer Science and Engineering",
			Year:        "2015 - 2019",
		},
	}
}

func CreateSkills() *Skills {
	return &Skills{
		Domains: []Skill{
			{
				ID:         uuid.NewString(),
				Title:      "Software Development",
				Percentage: 95,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Cloud Computing",
				Percentage: 90,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Cyber Security",
				Percentage: 85,
			},
			{
				ID:         uuid.NewString(),
				Title:      "UI/UX Design",
				Percentage: 80,
			},
		},
		Languages: []Skill{
			{
				ID:         uuid.NewString(),
				Title:      "Python",
				Percentage: 95,
			},
			{
				ID:         uuid.NewString(),
				Title:      "JavaScript",
				Percentage: 85,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Golang",
				Percentage: 80,
			},
		},
		Frameworks: []Skill{
			{
				ID:         uuid.NewString(),
				Title:      "Django",
				Percentage: 95,
			},
			{
				ID:         uuid.NewString(),
				Title:      "React",
				Percentage: 85,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Gin",
				Percentage: 70,
			},
		},
		Tools: []Skill{
			{
				ID:         uuid.NewString(),
				Title:      "Docker",
				Percentage: 90,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Git",
				Percentage: 90,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Linux",
				Percentage: 85,
			},
		},
		Services: []Skill{
			{
				ID:         uuid.NewString(),
				Title:      "AWS",
				Percentage: 90,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Firebase",
				Percentage: 85,
			},
			{
				ID:         uuid.NewString(),
				Title:      "Github Actions",
				Percentage: 100,
			},
		},
	}
}

func CreateMyServices() []MyService {
	return []MyService{
		{
			ID:          uuid.NewString(),
			Title:       "Backend Development",
			Description: "Build scalable and secure backend applications using stacks such as Django, Flask, FastAPI, express, Gin and Golang.",
			Icon:        "https://assets.satheesh.dev/icons/code-s-slash-line.svg",
		},
		{
			ID:          uuid.NewString(),
			Title:       "Frontend Development",
			Description: "Build responsive and interactive frontend applications using stacks such as React, Next.js, Bootstrap and Tailwind.",
			Icon:        "https://assets.satheesh.dev/icons/quill-pen-line.svg",
		},
		{
			ID:          uuid.NewString(),
			Title:       "Native App Development",
			Description: "Build cross-platform, user-friendly, modern and native applications using stacks such as Flutter, React Native and Electron.",
			Icon:        "https://assets.satheesh.dev/icons/smartphone-line.svg",
		},
	}
}

func CreateProjects() []Project {
	return []Project{
		{
			ID:         uuid.NewString(),
			Title:      "Code-Lordz",
			ShortDesc:  "Web application to help students learn and practice coding.",
			CoverImage: "https://miro.medium.com/max/1400/1*LXmzfblzOSgcQMmZnI3hNw.webp",
			Link:       "https://github.com/satheesh1997/node-codecompiler-ide",
			Filters: []string{
				"Node.js",
				"Express",
				"Bootstrap",
				"MongoDB",
			},
		},
		{
			ID:         uuid.NewString(),
			Title:      "Socket.io-Chat-Server",
			ShortDesc:  "Simple Chat Server Using Nodejs & Socket.io",
			CoverImage: "https://miro.medium.com/max/1400/1*J703pcFW-ezM87hNVPdM1Q.webp",
			Link:       "https://github.com/satheesh1997/Socket.io-Chat-Server",
			Filters: []string{
				"Node.js",
				"Socket.io",
			},
		},
		{
			ID:         uuid.NewString(),
			Title:      "iaBot",
			ShortDesc:  "Framework to build bots",
			CoverImage: "https://miro.medium.com/max/1400/1*RXEDFVThZhVfOmvSpG5LXw.webp",
			Link:       "https://github.com/satheesh1997/iaBot",
			Filters: []string{
				"Python",
			},
		},
		{
			ID:         uuid.NewString(),
			Title:      "Webpty",
			ShortDesc:  "Access shell based applications via browser",
			CoverImage: "https://miro.medium.com/max/1400/1*Tp0eLs1cvhLRRInbeDvzxA.webp",
			Link:       "https://github.com/satheesh1997/webpty",
			Filters: []string{
				"Python",
				"Tornado",
				"Xterm.js",
			},
		},
		{
			ID:         uuid.NewString(),
			Title:      "Bootstrap Login",
			ShortDesc:  "Simple login page UI with bootstrap",
			CoverImage: "https://raw.githubusercontent.com/satheesh1997/bootstrap-login/master/boot_login.png",
			Link:       "https://github.com/satheesh1997/bootstrap-login",
			Filters: []string{
				"HTML",
				"CSS",
				"Bootstrap",
			},
		},
		{
			ID:         uuid.NewString(),
			Title:      "PostHog",
			ShortDesc:  "Resolved an issue",
			CoverImage: "https://posthog-static-files.s3.us-east-2.amazonaws.com/Documentation-Assets/posthog-app-screenshot.png",
			Link:       "https://github.com/PostHog/posthog/pull/2461",
			Filters: []string{
				"Django",
				"Unit Testing",
				"Python",
			},
		},
	}
}

func CreateTestimonials() []Review {
	return []Review{
		{
			ID:           uuid.NewString(),
			Reviewer:     "Deepak Mishra",
			ReviewerMeta: "Manager - hackerearth.com",
			Review:       "Satheesh is dependable and his words can be trusted. He has the attitude of getting the work done and unblocking anyone dependent on him ASAP.",
			Rating:       4.7,
		},
		{
			ID:           uuid.NewString(),
			Reviewer:     "Ashutosh Sharma",
			ReviewerMeta: "Team Lead - hackerearth.com",
			Review:       "Satheesh is an integral part of our team in HackerEarth and he knows in and out about the product. There is always something to learn from him.",
			Rating:       5,
		},
		{
			ID:           uuid.NewString(),
			Reviewer:     "Chandransh Srivastava",
			ReviewerMeta: "UI Lead - hackerearth.com",
			Review:       "Satheesh's penchant for writing quality code within stipulated timelines makes him a great asset to any team. Someone, I would love to work again with.",
			Rating:       4.5,
		},
		{
			ID:           uuid.NewString(),
			Reviewer:     "Udit Narayan",
			ReviewerMeta: "Backend Lead - hackerearth.com",
			Review:       "I never had seen Satheesh shy away from a challenge. I wish Satheesh all the very best for his future.",
			Rating:       4.2,
		},
		{
			ID:           uuid.NewString(),
			Reviewer:     "Ashish Nasa",
			ReviewerMeta: "Product Designer - hackerearth.com",
			Review:       "Satheesh gives his 100% to every task that he is working on. In addition to being skilled at work, he is also a very helpful colleague and an awesome teammate.",
			Rating:       5,
		},
	}
}
