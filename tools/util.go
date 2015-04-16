package tools

import (
	"errors"
	"reflect"
	"time"
)

const (
	STORAGE_MAIN_MAP = iota
	STORAGE_READ_MAP
	STORAGE_RECENT_USER
	STORAGE_PERMANENT_EXIST
	STORAGE_CUSTOM_RULE
)

const (
	STORAGE_MAP_ADD = iota
	STORAGE_MAP_DEL
	STORAGE_MAP_UPD
)

const (
	DEFAULT_DURATION_TIME = time.Minute * 30
)

const (
	STORAGE_DEFAULT_SIZE         = 10000
	STORAGE_USAGE_AMOUNT float64 = 0.5
)

var (
	NilKeyError         = errors.New("nil key error")
	TimeOutError        = errors.New("the entity is die")
	HasEntError         = errors.New("old data is erased")
	NotEntError         = errors.New("not entity remove")
	NotEqualError       = errors.New("map[key] and value are not equal")
	RepeatNameError     = errors.New("add repeat name")
	StorageNotFindError = errors.New("can not find the name of storage list")
	ParameterTypeError  = errors.New("Parameter error")
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

/**
 * Check the map key type
 * @param val key
 * @return ok if key == nil || key type == Chan, Func, Map, Ptr, Interface, Slice return false
 * @return t type of err, 1: nil, 2: the map unsupport type
 */
func ChKeyType(val interface{}) (ok bool, t int) {
	if val == nil {
		return false, 1
	}
	rv := reflect.ValueOf(val)
	k := rv.Type().Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return rv.IsNil(), 2
	default:
		return true, 0
	}
}
