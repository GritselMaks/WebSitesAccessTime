package utils

import (
	"bufio"
	"net/url"
	"os"
)

// LoadUrlsList load url from file and validate
// If the scheme is not specified, 'http' will be added.
func LoadUrlsList(path string) ([]*url.URL, error) {
	listUrls, err := loadListURLFromFile(path)
	if err != nil {
		return nil, err
	}
	return MakeUrls(listUrls)
}

func loadListURLFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sitesList := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sitesList = append(sitesList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return sitesList, nil
}

func MakeUrls(list []string) ([]*url.URL, error) {
	urls := make([]*url.URL, 0, len(list))
	for _, s := range list {
		u, err := url.Parse(s)
		if err != nil {
			continue
		}
		if u.Scheme == "" {
			u.Scheme = "http"
		}
		if u.Host == "" {
			if u.Path != "" {
				u.Host = u.Path
				u.Path = ""
			} else {
				continue
			}
		}
		urls = append(urls, u)
	}
	return urls, nil
}
