package api

import (
	"github.com/MewX/KanTV-downloader-cli/kantv/util"
)

// GetCountryRequest store the request object for getting countries.
type GetCountryRequest struct {
	// No new fields, so left empty.
}

// NewGetCountryRequest creates the new request object.
func NewGetCountryRequest() Request {
	// Nothing to add.
	var req GetCountryRequest
	return Request{
		rtype:       POST,
		relativeURL: "Country/getCountry",
		body:        util.StructToURLValues(&req),
	}
}
