package api

import (
	"net/url"

	"github.com/MewX/KanTV-downloader-cli/kantv/util"
)

// GetCountryRequest store the request object for getting coutries.
type GetCountryRequest struct {
	// No new fields, so left empty.
}

// NewGetCountryRequest creates the new request object.
// The return values are: URL and request bytes.
// TODO: make all URL and request bytes in one struct.
func NewGetCountryRequest() (string, url.Values) {
	// Nothing to add.
	var req GetCountryRequest
	return "Country/getCountry", util.StructToURLValues(&req)
}
