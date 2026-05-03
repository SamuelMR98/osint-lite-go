package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

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

var sites = []Site{
	{"GitHub", "https://github.com/%s"},
	{"Reddit", "https://www.reddit.com/user/%s"},
	{"Hacker News", "https://news.ycombinator.com/user?id=%s"},
	{"DEV.to", "https://dev.to/%s"},
	{"Medium", "https://medium.com/@%s"},
	{"GitLab", "https://gitlab.com/%s"},
}

func checkSite(client *http.Client, site Site, username string) Result {
	url := fmt.Sprintf(site.URL, username)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Result{Site: site.Name, URL: url, Error: err.Error()}
	}

	req.Header.Set("User-Agent", "osint-lite-go/0.1")

	resp, err := client.Do(req)
	if err != nil {
		return Result{Site: site.Name, URL: url, Error: err.Error()}
	}
	defer resp.Body.Close()

	found := resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusMovedPermanently ||
		resp.StatusCode == http.StatusFound ||
		resp.StatusCode == http.StatusForbidden

	return Result{
		Site:       site.Name,
		URL:        url,
		Found:      found,
		StatusCode: resp.StatusCode,
	}
}

func main() {
	if len(os.Args) < 2 {
		color.Red("Usage: osint-lite <username>")
		os.Exit(1)
	}

	username := strings.TrimSpace(os.Args[1])
	// TODO: Implement usernames[] := usernameVariations(username) to generate common username variations (e.g., with dots, underscores, etc.)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	results := make(chan Result, len(sites))
	var wg sync.WaitGroup

	for _, site := range sites {
		wg.Add(1)

		go func(site Site) {
			defer wg.Done()
			results <- checkSite(client, site, username)
		}(site)
	}

	wg.Wait()
	close(results)

	color.Green("\nOSINT Lite Results for: %s\n\n", username)

	for result := range results {
		if result.Error != "" {
			color.Red("Error checking %s: %s\n", result.Site, result.Error)
		} else if result.Found {
			color.Green("[%s] %s (Status: %d)\n", result.Site, result.URL, result.StatusCode)
		} else {
			color.Yellow("[%s] %s (Status: %d)\n", result.Site, result.URL, result.StatusCode)
		}
	}
}