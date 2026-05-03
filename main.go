package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/SamuelMR98/osint-lite-go/internal"
	"github.com/fatih/color"
)


var sites = []internal.Site{
	{Name: "GitHub", URL: "https://github.com/%s"},
	{Name: "Reddit", URL: "https://www.reddit.com/user/%s"},
	{Name: "Hacker News", URL: "https://news.ycombinator.com/user?id=%s"},
	{Name: "DEV.to", URL: "https://dev.to/%s"},
	{Name: "Medium", URL: "https://medium.com/@%s"},
	{Name: "GitLab", URL: "https://gitlab.com/%s"},
}

func checkSite(client *http.Client, site internal.Site, username string) internal.Result {
	url := fmt.Sprintf(site.URL, username)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return internal.Result{Site: site.Name, URL: url, Error: err.Error()}
	}

	req.Header.Set("User-Agent", "osint-lite-go/0.1")

	resp, err := client.Do(req)
	if err != nil {
		return internal.Result{Site: site.Name, URL: url, Error: err.Error()}
	}
	defer resp.Body.Close()

	found := resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusMovedPermanently ||
		resp.StatusCode == http.StatusFound ||
		resp.StatusCode == http.StatusForbidden

	return internal.Result{
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

	results := make(chan internal.Result, len(sites))
	var wg sync.WaitGroup

	for _, site := range sites {
		wg.Add(1)

		go func(site internal.Site) {
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