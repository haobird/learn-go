package main

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

var (
	deviceLocalCache = gcache.New(1 << 12).LRU().Build()
)

type Device struct {
	DeviceId string
	NodeId   int64
	ParentId string
}

func main() {
	val := &Device{
		DeviceId: "dd",
		NodeId:   1,
	}
	var val2 *Device

	did1 := "a"
	did2 := "b"
	_ = deviceLocalCache.SetWithExpire(did1, val, time.Hour*24)
	_ = deviceLocalCache.SetWithExpire(did2, val2, time.Hour*24)
	result, err := deviceLocalCache.Get(did1)
	fmt.Printf("%+v, %v", result, err)

	result, err = deviceLocalCache.Get(did2)
	fmt.Printf("%+v, %v, %v", result, err, result == nil)
	result, err = deviceLocalCache.Get(did2)
	fmt.Printf("%+v, %v, %v", result.(*Device), err, result.(*Device) == nil)

}
