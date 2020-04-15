package api

// TODO: considering change the name of this file to account.go
type Login struct {
	Request

	noncestr         string // A random UUID
	timestamp        uint64 //
	sbID             string // appid
	sign             string // a hash result based on timestamp, noncestr, sbID, appid
	devicetype       string // always "0"
	interfaceVersion string // interface version
}
