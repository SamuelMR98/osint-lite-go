package internal

type Site struct {
	Name string
	URL  string
}

type Result struct {
	Site   string
	URL   string
	Found bool
	StatusCode int
	Error string
}
