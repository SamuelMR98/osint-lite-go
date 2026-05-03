package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/SamuelMR98/osint-lite-go/cmd"
	"github.com/SamuelMR98/osint-lite-go/internal"
	"github.com/SamuelMR98/osint-lite-go/utils"
	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)


var sites = []internal.Site{
	{Name: "GitHub", URL: "https://github.com/%s"},
	{Name: "Reddit", URL: "https://www.reddit.com/user/%s"},
	{Name: "Hacker News", URL: "https://news.ycombinator.com/user?id=%s"},
	{Name: "DEV.to", URL: "https://dev.to/%s"},
	{Name: "Medium", URL: "https://medium.com/@%s"},
	{Name: "GitLab", URL: "https://gitlab.com/%s"},
}

func main() {
	// Define flags with shorthand versions
	helpFlag := flag.BoolP("help", "h", false, "Show help message")
	flag.Parse()

	// Show help if the flag is set or if no username is provided
	if *helpFlag || flag.NArg() == 0 {
		cmd.Help()
		return
	}

	username := flag.Arg(0)
	color.Cyan("\nChecking availability for username: %s\n\n", username)

	client := &http.Client{Timeout: 10 * time.Second}
	var wg sync.WaitGroup
	results := make(chan internal.Result, len(sites))

	for _, site := range sites {
		wg.Add(1)
		go func(site internal.Site) {
			defer wg.Done()
			result := cmd.CheckSite(client, site, username)
			results <- result
		}(site)
	}

	wg.Wait()
	close(results)

	for result := range results {
		if result.Error != "" {
			color.Red("Error checking %s: %s\n", result.Site, result.Error)
		} else if result.Found {
			color.Green("%s: Found at %s (Status Code: %d - %s)\n", result.Site, result.URL, result.StatusCode, utils.GetStatusText(result.StatusCode))
		} else {
			color.Yellow("%s: Not found (Status Code: %d - %s)\n", result.Site, result.StatusCode, utils.GetStatusText(result.StatusCode))
		}
	}
}