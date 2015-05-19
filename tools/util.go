package tools

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"runtime/pprof"
	"time"
)

type STORAGE_MAP_TYPE int

const (
	STORAGE_MAIN_MAP        STORAGE_MAP_TYPE = 0
	STORAGE_READ_MAP        STORAGE_MAP_TYPE = 1
	STORAGE_RECENT_USER     STORAGE_MAP_TYPE = 2
	STORAGE_PERMANENT_EXIST STORAGE_MAP_TYPE = 3
	STORAGE_CUSTOM_RULE     STORAGE_MAP_TYPE = 4
)

type STORAGE_OPMAP_TYPE int

const (
	STORAGE_MAP_ADD STORAGE_OPMAP_TYPE = 0
	STORAGE_MAP_DEL STORAGE_OPMAP_TYPE = 1
	STORAGE_MAP_UPD STORAGE_OPMAP_TYPE = 2
)

const (
	DEFAULT_DURATION_TIME = time.Minute * 30
	DEFAULT_CLEARUP_TIME  = time.Minute * 5
	DEFAULT_STORAGE_TIME  = time.Minute * 10
)

const (
	STORAGE_DEFAULT_SIZE         = 10000
	STORAGE_USAGE_AMOUNT float64 = 0.5
	DEFAULT_BUFFER_SIZE          = 4096
)

const (
	SERIALIZATION_FILE_PATH = "../data/map"
	CONFIG_FILE_PATH        = "../conf/config.json"
)

type OpMapType string

const (
	TYPE_INT    OpMapType = "int"
	TYPE_INT32  OpMapType = "int"
	TYPE_INT64  OpMapType = "int"
	TYPE_BOOL   OpMapType = "int"
	TYPE_UINT   OpMapType = "int"
	TYPE_BYTE   OpMapType = "int"
	TYPE_STRING OpMapType = "int"
	TYPE_SLICE  OpMapType = "int"
	TYPE_MAP    OpMapType = "int"
	TYPE_SCRUCT OpMapType = "int"
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
	FileNotFouldError   = errors.New("Parameter error")
	FileEmplyError      = errors.New("Parameter error")
)

var (
	Debug            = true
	cpuProfile       = flag.String("cpuprofile", "", "write cpu profile to file")
	memProfile       = flag.String("memProfile", "", "write memory profile to file")
	blockProfile     = flag.String("blockProfile", "", "write block profile to file")
	cpuProfileRate   = flag.Int("cpuProfileRate", 0, "cpuProfileRate")
	memProfileRate   = flag.Int("memProfileRate", 0, "memProfileRate")
	blockProfileRate = flag.Int("blockProfileRate", 0, "blockProfileRate")
)

type NetOpType string

const (
	NET_TYPE_GET NetOpType = "get"
	NET_TYPE_PUT NetOpType = "put"
)

type MyServerFlag int

const (
	TCP_CLIENT MyServerFlag = 0
	TCP_SERVER MyServerFlag = 1
)

type ProfileType string

const (
	RPOFILE_HEAP ProfileType = "heap"
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

func ChType(val interface{}) (t string) {
	if val == nil {
		return "nil"
	}
	switch val.(type) {
	case int:
		return "int"
	default:
		return "other"
	}
}

func Printf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		return fmt.Printf(format, a...)
	} else {
		return 0, nil
	}
}

func Println(a ...interface{}) (n int, err error) {
	if Debug {
		return fmt.Println(a...)
	} else {
		return 0, nil
	}
}

func WhatType(i interface{}) {
	fmt.Printf("%T\n", i)
}

func StartCPUProfile() {
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return
		}
	}
}

func StopCPUProfile() {
	if *cpuProfile != "" {
		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
	}
}

func StartMemProfile() {
	if *memProfile != "" && *memProfileRate > 0 {
		runtime.MemProfileRate = *memProfileRate
	}
}

func StopMemProfile() {
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create mem profile output file: %s", err)
			return
		}
		if err = pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", *memProfile, err)
		}
		f.Close()
	}
}

func StartBlockProfile() {
	if *blockProfile != "" && *blockProfileRate > 0 {
		runtime.SetBlockProfileRate(*blockProfileRate)
	}
}

func StopBlockProfile() {
	if *blockProfile != "" && *blockProfileRate >= 0 {
		f, err := os.Create(*blockProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create block profile output file: %s", err)
			return
		}
		if err = pprof.Lookup("block").WriteTo(f, 0); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", *blockProfile, err)
		}
		f.Close()
	}
}

func SaveProfile(workDir string, profileName string, ptype ProfileType, debug int) {
	if profileName == "" {
		profileName = string(ptype)
	}
	profilePath := filepath.Join(workDir, profileName)
	f, err := os.Create(profilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not create profile output file: %s", err)
		return
	}
	if err = pprof.Lookup(string(ptype)).WriteTo(f, debug); err != nil {
		fmt.Fprintf(os.Stderr, "Can not write %s: %s", profilePath, err)
	}
	f.Close()
}

func ChErr(err error) {
	if err != nil {
		Println(err)
	}
}
