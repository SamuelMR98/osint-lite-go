package internal

type Site struct {
	Name string
	URL  string
}

type Result struct {
	Site       string
	URL        string
	Found      bool
	StatusCode int
	Error      string
}
type CheckOptions struct {
	Social bool
	Tech   bool

	PrintJSON bool
	SaveJSON  string
	NoSpinner bool
}

var SocialSites = []Site{
	{Name: "Instagram", URL: "https://www.instagram.com/%s"},
	{Name: "Twitter", URL: "https://twitter.com/%s"},
	{Name: "Facebook", URL: "https://www.facebook.com/%s"},
	{Name: "LinkedIn", URL: "https://www.linkedin.com/in/%s"},
	{Name: "TikTok", URL: "https://www.tiktok.com/@%s"},
	{Name: "Reddit", URL: "https://www.reddit.com/user/%s"},
}

var TechSites = []Site{
	{Name: "GitHub", URL: "https://github.com/%s"},
	{Name: "Hacker News", URL: "https://news.ycombinator.com/user?id=%s"},
	{Name: "DEV.to", URL: "https://dev.to/%s"},
	{Name: "Medium", URL: "https://medium.com/@%s"},
	{Name: "GitLab", URL: "https://gitlab.com/%s"},
}

func GetSocialSites() []Site {
	return SocialSites
}

func GetTechSites() []Site {
	return TechSites
}
