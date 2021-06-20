package window

import "sync/atomic"

type Bucket struct {
	// 标识是哪一秒的桶数据
	windowStart int64
	// 用于简单自增统计数据
	adder []int32
	// 最大并发类的统计数据
	maxUpdater []int32
}

func NewBucket(windowStart int64) *Bucket {

	adder := make([]int32, EventCount)
	maxUpdater := make([]int32, EventCount)
	return &Bucket{
		windowStart: windowStart,
		adder:       adder,
		maxUpdater:  maxUpdater,
	}
}

func (b *Bucket) GetAdder(event Event) int32 {

	return b.adder[event]

}

func (b *Bucket) GetMaxUpdater(event Event) int32 {

	return b.maxUpdater[event]

}

func (b *Bucket) Increment(event Event) {
	atomic.AddInt32(&b.adder[event],1)

}

func (b *Bucket) Add(event Event, value int32) {
	atomic.AddInt32(&b.adder[event],value)

}

func (b *Bucket) UpdateMaxUpdater(event Event, value int32) {
	if b.maxUpdater[event] < value {
		atomic.StoreInt32(&b.maxUpdater[event],value)
	}

}
