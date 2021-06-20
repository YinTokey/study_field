package main

import (
	. "Week06/window"
	"fmt"
	"time"
)

func main() {
	r := NewRollingNumber(2,5)

	time.Sleep(time.Second * 1)
	r.Increment(EVENT_SUCCESS)
	time.Sleep(time.Second * 1)
	fmt.Printf("bucket 1 %+v \n", r.GetCurrentBucket())
	r.Increment(EVENT_FAILURE)
	time.Sleep(time.Second * 1)
	r.Increment(EVENT_FAILURE)
	time.Sleep(time.Second * 1)
	r.Increment(EVENT_FAILURE)
	time.Sleep(time.Second * 1)
	r.Increment(EVENT_TIMEOUT)
	time.Sleep(time.Second * 1)
	r.Increment(EVENT_REJECTION)
	time.Sleep(time.Second * 1)
	r.UpdateRollingMax(EVENT_REJECTION, 5)
	time.Sleep(time.Second * 1)
	fmt.Printf("bucket 2 %+v \n", r.GetCurrentBucket())

}
