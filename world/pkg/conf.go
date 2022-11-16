package pkg

import (
	"reflect"
)

func LoadConf(filepath string, out interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return nil

}
