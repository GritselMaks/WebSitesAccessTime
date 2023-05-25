package utils

import (
	"os"
	"testing"
)

func TestValidateUrls(t *testing.T) {
	urls := []string{"https://example.com", "http://example.com", "example.com", "192.168.210.111:8080", "http://"}
	validatedUrls, err := MakeUrls(urls)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expectedSchemes := []string{"https", "http", "http"}
	for i, u := range validatedUrls {
		if u.Scheme != expectedSchemes[i] {
			t.Errorf("expected '%s' for URL %d, get: '%s'", expectedSchemes[i], i, u.Scheme)
		}
	}
}

func TestLoadUrlsList(t *testing.T) {
	//create file
	testPath := "./test.txt"
	f, err := os.Create(testPath)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testPath)
	defer f.Close()
	b := []byte(`https://example.com
https://google.com
https://stackoverflow.com`)
	_, err = f.Write(b)
	if err != nil {
		t.Fatal(err)
	}

	got, err := LoadUrlsList(testPath)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	expectCountOfURL := 3
	if len(got) != expectCountOfURL {
		t.Errorf("expected '%v'URLs, get: '%v'", expectCountOfURL, len(got))
	}

	//Invalid path
	_, err = LoadUrlsList("testPath")
	if err == nil {
		t.Errorf("Expected err, but got nil")
	}

}
