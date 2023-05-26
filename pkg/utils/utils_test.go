package utils

import (
	"os"
	"testing"
)

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

func TestMakeUrl(t *testing.T) {
	urls := []string{"https://example.com", "http://example.com", "example.com", "http://192.168.210.111:8080", "http://"}
	validatedUrls := makeUrlList(urls)
	if len(validatedUrls) != len(urls)-1 {
		t.Errorf("expected [%v], get: [%v]\n", len(urls)-1, len(validatedUrls))
	}
}
