package me

type (
	Controller struct{}
	Service    struct{}
)

type (
	Skill struct {
		ID         string `json:"id"`
		Title      string `json:"title"`
		Percentage int    `json:"percentage"`
	}

	// The name of the struct is kept as MyService to avoid conflict with the Service struct
	// in the same package
	MyService struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}

	Project struct {
		ID         string   `json:"id"`
		Title      string   `json:"title"`
		ShortDesc  string   `json:"shortDesc"`
		CoverImage string   `json:"coverImage"`
		Link       string   `json:"link"`
		Filters    []string `json:"filters"`
		Images     []string `json:"images"`
		Videos     []string `json:"videos"`
	}

	Review struct {
		ID            string  `json:"id"`
		Reviewer      string  `json:"reviewer"`
		ReviewerMeta  string  `json:"reviewerMeta"`
		ReviewerImage string  `json:"reviewerImage"`
		Review        string  `json:"review"`
		Rating        float64 `json:"rating"`
	}
)

type (
	Info struct {
		FirstName      string   `json:"firstName"`
		LastName       string   `json:"lastName"`
		FullName       string   `json:"fullName"`
		Bio            string   `json:"bio"`
		About          string   `json:"about"`
		DateOfBirth    string   `json:"dob"`
		Address        string   `json:"address"`
		ImageURL       string   `json:"image"`
		ThumbnailURL   string   `json:"thumbImage"`
		Languages      []string `json:"languages"`
		PhoneNumbers   []string `json:"phoneNumbers"`
		EmailAddresses []string `json:"emailAddresses"`
		SocialProfile  struct {
			Whatsapp string `json:"whatsapp"`
			LinkedIn string `json:"linkedin"`
			Twitter  string `json:"twitter"`
			Github   string `json:"github"`
			Medium   string `json:"medium"`
		} `json:"socialProfile"`
	}

	Education struct {
		ID          string `json:"id"`
		Institution string `json:"institution"`
		Degree      string `json:"degree"`
		Details     string `json:"details"`
		Year        string `json:"year"`
	}

	Skills struct {
		Domains    []Skill `json:"domains"`
		Languages  []Skill `json:"programmingLanguages"`
		Frameworks []Skill `json:"frameworks"`
		Tools      []Skill `json:"tools"`
		Services   []Skill `json:"services"`
	}
)
