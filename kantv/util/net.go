package util

import (
	"io/ioutil"
	"net/http"
)

// FetchLinkContent downloads the link content to variable in memory.
func FetchLinkContent(url string) ([]byte, error) {
	// Just get it!
	resp, errNet := http.Get(url) // No need to set User-Agent.
	if errNet != nil {
		return []byte(""), errNet
	}
	defer resp.Body.Close()

	content, errDecode := ioutil.ReadAll(resp.Body)
	if errDecode != nil {
		return []byte(""), errDecode
	}

	return content, nil
}
