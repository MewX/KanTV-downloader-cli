package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	urllib "net/url"
	"strconv"
	"strings"

	"github.com/MewX/KanTV-downloader-cli/kantv/util"
)

// Request is the 'generic' type for all network requests.
type Request interface {
	Decode(b []byte)
	Encode(r Request) []byte
}

// TODO
//func (br Request) Decode(b []byte) {
//
//}
//
//func (br Request) Encode(o *BaseRequest) {
//
//}

// SendRequest sends the request to the API server.
// TODO: should automatically use fallback domains.
func SendRequest(url string, request urllib.Values) (map[string]interface{}, error) {
	// session := &http.Client{Transport: &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }}

	// Generate a new sign.
	s := NewSign()
	for k, v := range util.StructToURLValues(&s) {
		request[k] = v
	}
	urlEncodedData := request.Encode()

	client := &http.Client{}
	// TODO: should support POST/GET.
	req, _ := http.NewRequest("POST", getDomain()+url, strings.NewReader(urlEncodedData))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(urlEncodedData)))

	// Make HTTP(S) request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Print the received json string when verbose is specified.
	if util.VerboseMode {
		var buf bytes.Buffer
		json.Indent(&buf, []byte(string(body)), "", "  ")
		fmt.Println(buf.String())
	}

	// Decode the json string to map.
	var obj map[string]interface{}
	json.Unmarshal([]byte(string(body)), &obj)
	return obj, err
}

// TODO: add test and make this generic
func getDomain() string {
	return "https://" + Host1 + "/"
}
