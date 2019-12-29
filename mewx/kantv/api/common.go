package api

const interface_version = "104"
const appid = "kantvea2663ea5d5evje09j"

type Sign struct {
	appid string
	noncestr string  // UUID
	sbID string  // device ID, this should be settable
	timestamp string  // a timestamp string
	sign string  // the hash
}

