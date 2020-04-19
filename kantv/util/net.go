package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// FetchLinkContentWithRetry downloads the link content to variable in memory with retry.
func FetchLinkContentWithRetry(url string) ([]byte, error) {
	var b []byte
	var e error
	for i := 0; i < RetryTimes; i++ {
		b, e = FetchLinkContent(url)
		if e == nil {
			return b, e
		}
	}
	return b, e
}

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

// ExtractM3u8BaseURL extracts the base URL for *.ts videos.
func ExtractM3u8BaseURL(url string) (string, error) {
	cutPoint := strings.Index(url, ".m3u8")
	if cutPoint == -1 {
		return "", fmt.Errorf("bad m3u8 URL: couldn't find m3u8 keyword in " + url)
	}

	cutLeftString := url[:cutPoint]
	endSlashPoint := strings.LastIndex(cutLeftString, "/")
	if endSlashPoint == -1 {
		return "", fmt.Errorf("bad m3u8 URL: couldn't find end slash in " + url)
	}

	baseURL := cutLeftString[:endSlashPoint+1]
	if VerboseMode {
		fmt.Printf("Extracted base URL is %s\n", baseURL)
	}
	return baseURL, nil
}

// ExtractTvidPartidFromURL extracts the IDs.
func ExtractTvidPartidFromURL(url string) (string, string, error) {
	// Try one without part ID.
	re := regexp.MustCompile(".+?/(tvdrama|anime|movie)/(\\d+)")
	match := re.FindStringSubmatch(url)
	if len(match) != 3 {
		return "", "", fmt.Errorf("url is not supported; if you believe this is an error, pleaser file an issue")
	}
	tvid := match[2]

	// Try one with part ID.
	re = regexp.MustCompile(".+?/(tvdrama|anime|movie)/(\\d+?)-(\\d+)")
	match = re.FindStringSubmatch(url)
	partid := ""
	if len(match) == 4 {
		partid = match[3]
	}
	return tvid, partid, nil
}
