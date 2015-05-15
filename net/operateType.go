package net

import (
	"encoding/json"
	"github.com/yanjinzh6/flowkey/tools"
	"time"
)

type OperateType struct {
	T   tools.NetOpType
	Key interface{}
	Val interface{}
	D   time.Duration
	Err error
}

type ResOperateType struct {
	T tools.NetOpType
}

func NewOperateType() *OperateType {
	return &OperateType{}
}

func (o *OperateType) Get(key interface{}) {
	o.T = tools.NET_TYPE_GET
	o.Key = key
}

func (o *OperateType) Put(key, val interface{}, dur time.Duration) {
	o.T = tools.NET_TYPE_PUT
	o.Key = key
	o.Val = val
	o.D = dur
}

func (o *OperateType) EncodeByJson() (data []byte, err error) {
	return json.Marshal(o)
}

func (o *OperateType) DecodeByJson(data []byte) (err error) {
	return json.Unmarshal(data, o)
}
