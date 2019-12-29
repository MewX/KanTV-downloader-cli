package api

type Request interface {
	Decode(b []byte)
	Encode(r Request) []byte
}

//func (br Request) Decode(b []byte) {
//
//}
//
//func (br Request) Encode(o *BaseRequest) {
//
//}