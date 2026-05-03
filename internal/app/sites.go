package app

// This file contains the main logic for checking username availability across various platforms.
// It defines the list of platforms to check, performs HTTP requests to determine if the username exists,
// and formats the output for the user.

import (
	"fmt"
	"net/http"

	"github.com/SamuelMR98/osint-lite-go/internal"
)

func CheckSite(client *http.Client, site internal.Site, username string) internal.Result {
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

func CheckSites(username string, sites []internal.Site) []internal.Result {
	client := &http.Client{}
	results := make([]internal.Result, len(sites))

	for i, site := range sites {
		results[i] = CheckSite(client, site, username)
	}

	return results
}

func BuildSelectedSites(opts internal.CheckOptions) []internal.Site {
	var selectedSites []internal.Site

	if opts.Social {
		selectedSites = append(selectedSites, internal.GetSocialSites()...)
	}
	if opts.Tech {
		selectedSites = append(selectedSites, internal.GetTechSites()...)
	}

	// Default behaviour: run all categories if none specified
	if !opts.Social && !opts.Tech {
		selectedSites = append(selectedSites, internal.GetSocialSites()...)
		selectedSites = append(selectedSites, internal.GetTechSites()...)
	}

	return selectedSites
}