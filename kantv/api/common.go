package api

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
)

// Consts from KanTV apk.
const (
	// com.kantv.common.utils.Utils
	InterfaceVerion = "106"

	// com.kantv.ui.config.SignConfig
	AppID = "kantvea2663ea5d5evje09j"

	// com.kantv.common.BuildConfig
	ApplicationID = "com.kantv.common"
	Host1         = "www.suramic.com"
	Host2         = "skyollie.com"

	// com.kantv.common.api.Api
	UserAgent = "suramic app1.0"
)

// Sign contains the shard fields for every network request.
type Sign struct {
	appid     string
	noncestr  string // UUID
	sbID      string // device ID, this should be settable
	timestamp string // a timestamp string
	sign      string // the hash

	// Other fields that are not related to signatures.
	devicetype      string // always "0"
	interfaceVerion string // the API interface version
}

// NewSign makes a new sign object. Reference: com.kantv.ui.config.SignConfig.
func NewSign() Sign {
	var s Sign
	s.appid = AppID
	s.noncestr = "bb28a26e-2867-43e9-a1ce-ba89897efd71" // TODO(#18): use random UUID.
	s.sbID = "977684d6d549c2e5"                         // TODO(#19): Allow user to set the random device ID.
	s.timestamp = strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	// Calculate the sign.
	toBeHashed := fmt.Sprintf("noncestr=%s&sbID=%s&timestamp=%s&appid=%s",
		s.noncestr, s.sbID, s.timestamp, s.appid)
	s.sign = fmt.Sprintf("%x", sha1.Sum([]byte(toBeHashed)))

	// Other fields.
	s.devicetype = "0"
	s.interfaceVerion = InterfaceVerion

	return s
}
