package api

import (
	"github.com/MewX/KanTV-downloader-cli/kantv/util"
)

// GetVideoInfoRequest stores the request object for getting video detailed info.
type GetVideoInfoRequest struct {
	_token string
	tvid   string
	// This one disobeys the goling rules, but I had to say KanTV developers suck.
	part_id string // Can be empty for requesting full information.
}

// NewGetVideoInfoRequest creates the new request object.
func NewGetVideoInfoRequest(tvid string) Request {
	req := GetVideoInfoRequest{
		_token:  util.UserToken,
		tvid:    tvid,
		part_id: "",
	}
	return Request{
		rtype:       POST,
		relativeURL: "Index/play",
		body:        util.StructToURLValues(&req),
	}
}
