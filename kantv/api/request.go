package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	urllib "net/url"

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
// TODO: should support POST/GET.
// TODO: should automatically use fallback domains.
func SendRequest(url string, request urllib.Values) {
	// session := &http.Client{Transport: &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }}

	s := NewSign()
	for k, v := range util.StructToURLValues(&s) {
		request[k] = v
	}

	response, err := http.PostForm(getDomain()+url, request)
	if err != nil {
		// TODO: handle postform error
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// TODO: handle read response error
	}

	fmt.Printf("%s\n", string(body))
}

// TODO: add test and make this generic
func getDomain() string {
	return "https://" + Host1 + "/"
}
