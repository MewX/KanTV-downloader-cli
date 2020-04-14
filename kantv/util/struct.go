package util

import (
	"fmt"
	"net/url"
	"reflect"
)

// StructToURLValues convert a struct to url.Values map.
func StructToURLValues(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		values.Set(typ.Field(i).Name, fmt.Sprint(iVal.Field(i)))
	}
	return
}
