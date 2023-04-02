package me

import (
    "github.com/gin-gonic/gin"
)

func (controller Controller) GetInfo (c *gin.Context) {
    c.JSON(200, gin.H{
        "firstName": "Satheesh",
        "lastName": "Kumar",
        "fullName": "Satheesh Kumar",
        "thumbImage": "https://assets.satheesh.dev/images/formal-image.jpg",
        "largeImage": "https://assets.satheesh.dev/images/about-image.jpg",
        "bio": "Engineer @HackerEarth ||  Open Source Contributor  || Designer  ||  Developer || Mentor",
        "dob": "10-10-1997",
        "nationality": "Indian",
        "languages": []string{
            "English",
            "Tamil",
            "Hindi",
        },
        "address": "Bengaluru, Karnataka, India",
        "socialFootprints": gin.H{
            "github": "https://github.com/satheesh1997",
            "linkedin": "https://www.linkedin.com/in/satheesh1997/",
            "twitter": "https://twitter.com/git_push_tweet",
            "whatsapp": "https://wa.me/919597264229",
            "medium": "https://esc-wq.medium.com",
        },
        "phoneNumbers": []string{
            "+91-9597264229",
        },
        "emailAddresses": []string{
            "mail@satheesh.dev",
        },
    })
}

func (controller Controller) GetSkills (c *gin.Context) {
    c.JSON(200, gin.H{
        "domains": []gin.H{
            gin.H{
                "id": 1,
                "title": "Software Development",
                "percentage": 95,
            },
            gin.H{
                "id": 2,
                "title": "Cloud Computing",
                "percentage": 90,
            },
            gin.H{
                "id": 3,
                "title": "Cyber Security",
                "percentage": 85,
            },
            gin.H{
                "id": 4,
                "title": "UI/UX Design",
                "percentage": 80,
            },
        },
        "languages": []gin.H{
            gin.H{
                "id": 1,
                "title": "Python",
                "percentage": 95,
            },
            gin.H{
                "id": 2,
                "title": "JavaScript",
                "percentage": 85,
            },
            gin.H{
                "id": 3,
                "title": "Golang",
                "percentage": 80,
            },
        },
        "frameworks": []gin.H{
            gin.H{
                "id": 1,
                "title": "Django",
                "percentage": 95,
            },
            gin.H{
                "id": 2,
                "title": "React",
                "percentage": 85,
            },
            gin.H{
                "id": 3,
                "title": "Gin",
                "percentage": 70,
            },
        },
        "tools": []gin.H{
            gin.H{
                "id": 1,
                "title": "Docker",
                "percentage": 90,
            },
            gin.H{
                "id": 2,
                "title": "Git",
                "percentage": 95,
            },
            gin.H{
                "id": 3,
                "title": "Linux",
                "percentage": 90,
            },
        },
        "services": []gin.H{
            gin.H{
                "id": 1,
                "title": "AWS",
                "percentage": 90,
            },
            gin.H{
                "id": 2,
                "title": "Firebase",
                "percentage": 90,
            },
            gin.H{
                "id": 3,
                "title": "GitHub Actions",
                "percentage": 95,
            },
        },
    })
}

func (controller Controller) GetServices (c *gin.Context) {
    c.JSON(200, []gin.H{
        gin.H{
            "id": 1,
            "title": "Frontend",
            "description": "Build client-side applications with modern features like SPA and SEO optimization using stacks such as Next.js and tailwindCSS",
            "icon": "https://assets.satheesh.dev/icons/quill-pen-line.svg",
        },
        gin.H{
            "id": 2,
            "title": "Backend",
            "description": "Build scalable and secure backend applications using stacks such as Django, Flask, FastAPI, express, Gin and Golang",
            "icon": "https://assets.satheesh.dev/icons/code-s-slash-line.svg",
        },
        gin.H{
            "id": 3,
            "title": "Native Apps",
            "description": "Build cross-platform, user-friendly, modern and native applications using stacks such as Flutter, React Native and Electron",
            "icon": "https://assets.satheesh.dev/icons/smartphone-line.svg",
        },
    })
}

func (controller Controller) GetProjects (c *gin.Context) {
    c.JSON(200, []gin.H{
        gin.H{
            "id": 1,
            "title": "Code-Lordz",
            "subtitle": "Web-Based code execution tool",
            "coverimage": "https://miro.medium.com/max/1400/1*LXmzfblzOSgcQMmZnI3hNw.webp",
            "imagegallery": false,
            "videogallery": false,
            "url": "https://github.com/satheesh1997/node-codecompiler-ide",
            "filters": []string{
                "app",
                "UI",
            },
        },
        gin.H{
            "id": 2,
            "title": "Socket.io-Chat-Server",
            "subtitle": "Simple Chat Server Using Nodejs & Socket.io",
            "coverimage": "https://miro.medium.com/max/1400/1*J703pcFW-ezM87hNVPdM1Q.webp",
            "imagegallery": []string{
                "https://miro.medium.com/max/1400/1*J703pcFW-ezM87hNVPdM1Q.webp",
                "https://miro.medium.com/max/1400/1*dY5177srJGfAtWfUlMxLDg.webp",
            },
            "videogallery": false,
            "url": "https://github.com/satheesh1997/Socket.io-Chat-Server",
            "filters": []string{
                "app",
                "cli",
            },
        },
        gin.H{
            "id": 4,
            "title": "iaBot",
            "subtitle": "Framework to build bots",
            "coverimage": "https://miro.medium.com/max/1400/1*RXEDFVThZhVfOmvSpG5LXw.webp",
            "imagegallery": false,
            "videogallery": false,
            "url": "https://github.com/satheesh1997/iaBot",
            "filters": []string{
                "framework",
            },
        },
        gin.H{
            "id": 5,
            "title": "Webpty",
            "subtitle": "Access shell based applications via browser",
            "coverimage": "https://miro.medium.com/max/1400/1*Tp0eLs1cvhLRRInbeDvzxA.webp",
            "imagegallery": []string{
                "https://miro.medium.com/max/1400/1*Tp0eLs1cvhLRRInbeDvzxA.webp",
                "https://miro.medium.com/max/1400/0*X3sUVypk5G836ccP.webp",
                "https://miro.medium.com/max/1400/0*_fUACCnidEefSPvZ.webp",
                "https://miro.medium.com/max/1400/0*LdsHUuRPP6pkFQNK.webp",
            },
            "videogallery": false,
            "url": "https://github.com/PostHog/posthog/pull/2461",
            "filters": []string{
                "app",
                "cli",
            },
        },
        gin.H{
            "id": 6,
            "title": "Satheesh Kumar",
            "subtitle": "Developer portfolio website",
            "coverimage": "/images/portfolios/portfolio-image-1.png",
            "imagegallery": false,
            "videogallery": false,
            "url": "https://github.com/PostHog/posthog/pull/2461",
            "filters": []string{
                "UI",
            },
        },
        gin.H{
            "id": 3,
            "title": "Bootstrap Login",
            "subtitle": "Simple login page UI with bootstrap",
            "coverimage": "https://raw.githubusercontent.com/satheesh1997/bootstrap-login/master/boot_login.png",
            "imagegallery": false,
            "videogallery": false,
            "url": "https://github.com/satheesh1997/bootstrap-login",
            "filters": []string{
                "UI",
            },
        },
        gin.H{
            "id": 6,
            "title": "PostHog",
            "subtitle": "Resolved an issue",
            "coverimage": "https://posthog-static-files.s3.us-east-2.amazonaws.com/Documentation-Assets/posthog-app-screenshot.png",
            "imagegallery": false,
            "videogallery": false,
            "url": "https://github.com/PostHog/posthog/pull/2461",
            "filters": []string{
                "help",
            },
        },
    })
}

func (controller Controller) GetEducation (c *gin.Context) {
    c.JSON(200, []gin.H{
        gin.H{
            "id": 1,
            "degree": "B.E",
            "details": "Computer Science and Engineering",
            "institution": "Sri Krishna College of Engineering and Technology",
            "year": "2015-2019",
        },
        gin.H{
            "id": 2,
            "degree": "HSC",
            "details": "Computer Science, Mathematics, Physics, Chemistry",
            "institution": "Amrita Vidyalayam",
            "year": "2013-2015",
        },
    })
}

// Note: This will be replaced with linkedin api to fetch the data
// from next version
func (controller Controller) GetExperience (c *gin.Context) {
    c.JSON(200, []gin.H{
        gin.H{
            "id": 1,
            "title": "Software Engineer",
            "type": "Internship",
            "company": "HackerEarth",
            "companyLocation": "Bangalore, India",
            "date": gin.H{
                "start": "May 2018",
                "end": "Feb 2019",
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
            "id": 2,
            "title": "Software Engineer",
            "type": "Full Time",
            "company": "HackerEarth",
            "companyLocation": "Bangalore, India",
            "date": gin.H{
                "start": "Apr 2019",
                "end": "Feb 2021",
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
            "id": 3,
            "title": "Senior Software Engineer",
            "type": "Full Time",
            "company": "HackerEarth",
            "companyLocation": "Bangalore, India",
            "date": gin.H{
                "start": "Feb 2021",
                "end": "Apr 2022",
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
            "id": 4,
            "title": "Staff Software Engineer",
            "type": "Full Time",
            "company": "HackerEarth",
            "companyLocation": "Bangalore, India",
            "date": gin.H{
                "start": "Apr 2022",
                "end": "Present",
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
func (controller Controller) GetReviews (c *gin.Context) {
    c.JSON(200, []gin.H{
        gin.H{
            "id": 1,
            "name": "Deepak Mishra",
            "meta": "Manager - hackerearth.com",
            "rating": 4.7,
            "image": "https://media.licdn.com/dms/image/D4D35AQE5MGgOk4jFUA/profile-framedphoto-shrink_400_400/0/1657306604641?e=1671606000&v=beta&t=Ua14FIWszyeYfz74RrtwrqJyQt5PgDZDUG9oHx97dfM",
            "text": "Satheesh is dependable and his words can be trusted. He has the attitude of getting the work done and unblocking anyone dependent on him ASAP.",
        },
        gin.H{
            "id": 2,
            "name": "Ashutosh Sharma",
            "meta": "Team Lead - hackerearth.com",
            "rating": 5,
            "image": "https://media.licdn.com/dms/image/C5603AQH3qCL4yxtqhQ/profile-displayphoto-shrink_400_400/0/1517540083233?e=1676505600&v=beta&t=bvKRd-Iik4U97wkIEdYokDvMCgDDPz1_0uGBGOuAREs",
            "text": "Satheesh is an integral part of our team in HackerEarth and he knows in and out about the product. There is always something to learn from him.",
        },
        gin.H{
            "id": 3,
            "name": "Chandransh Srivastava",
            "meta": "UI Lead - hackerearth.com",
            "rating": 4.5,
            "image": "https://media.licdn.com/dms/image/C5603AQE0BUE9Ypnu9A/profile-displayphoto-shrink_400_400/0/1608583550420?e=1676505600&v=beta&t=PLSBTIFv3WK4u-97o4-I6UwEOr9jSMwz8yddzNIbjmU",
            "text": "Satheesh's penchant for writing quality code within stipulated timelines makes him a great asset to any team. Someone, I would love to work again with.",
        },
        gin.H{
            "id": 4,
            "name": "Udit Narayan",
            "meta": "Backend Lead - hackerearth.com",
            "rating": 4.2,
            "image": "https://media.licdn.com/dms/image/C5103AQFQVwzIRUawPg/profile-displayphoto-shrink_400_400/0/1586181266294?e=1676505600&v=beta&t=PyZE2hTfHi6JT0r9JUEa8K15zb-2EVsTCZajqVdPvA0",
            "text": "I never had seen Satheesh shy away from a challenge. I wish Satheesh all the very best for his future.",
        },
        gin.H{
            "id": 5,
            "name": "Ashish Nasa",
            "meta": "Product Designer - hackerearth.com",
            "rating": 5,
            "image": "https://media.licdn.com/dms/image/C4E03AQH7S6yVBSIz8Q/profile-displayphoto-shrink_400_400/0/1612773440098?e=1676505600&v=beta&t=CR7p7QWAHyw3mDkCD2k70J8Xmn80yQvUq8iYrAoz_V8",
            "text": "Satheesh gives his 100% to every task that he is working on. In addition to being skilled at work, he is also a very helpful colleague and an awesome teammate.",
        },
    })
}
