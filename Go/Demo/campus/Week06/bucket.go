package main

type Bucket struct {
	// 标识是哪一秒的桶数据
	windowStart int
	// 用于简单自增统计数据
	adder []int
	// 最大并发类的统计数据
	maxUpdater []int
}

func NewBucket(windowStart int) *Bucket {

	adder := make([]int, EventCount)
	maxUpdater := make([]int, EventCount)
	return &Bucket{
		windowStart: windowStart,
		adder:       adder,
		maxUpdater:  maxUpdater,
	}
}

func (b *Bucket) GetAdder(event Event) int {

	return b.adder[event]

}

func (b *Bucket) GetMaxUpdater(event Event) int {

	return b.maxUpdater[event]

}

func (b *Bucket) increment(event Event) {
	b.adder[event]++
}

func (b *Bucket) add(event Event, value int) {
	b.adder[event] += value
}

func (b *Bucket) updateMaxUpdater(event Event, value int) {
	if b.maxUpdater[event] < value {
		b.maxUpdater[event] = value
	}
}
