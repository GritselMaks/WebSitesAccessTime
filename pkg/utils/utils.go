package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

// LoadUrlsList load url from file and validate
// If the scheme is not specified, 'http' will be added.
func LoadUrlsList(path string) ([]string, error) {
	listUrls, err := loadListURLFromFile(path)
	if err != nil {
		return nil, err
	}
	return makeUrlList(listUrls), nil
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

func makeUrlList(list []string) []string {
	urls := make([]string, 0, len(list))
	for _, s := range list {
		url, err := MakeUrl(s)
		if err != nil {
			continue
		}
		urls = append(urls, url)
	}
	return urls
}

// MakeUrl return string with url in valid format or error
// Default scheme is http
func MakeUrl(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	if u.Host == "" && u.Path == "" {
		return "", fmt.Errorf("url is not valid")
	}
	return u.String(), nil
}
