package tools

import (
	"reflect"
)

/**
 * Check the map key type
 * @param val key
 * @return ok if key == nil || key type == Chan, Func, Map, Ptr, Interface, Slice return false
 */
func ChKey(val interface{}) (ok bool) {
	if val == nil {
		return false
	}
	rv := reflect.ValueOf(val)
	k := rv.Type().Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return rv.IsNil()
	default:
		return true
	}
}
