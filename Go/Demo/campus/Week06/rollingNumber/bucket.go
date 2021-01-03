package rollingNumber

type Bucket struct {
	// 标识是哪一秒的桶数据
	windowStart int64
	// 用于简单自增统计数据
	adder []int64
	// 最大并发类的统计数据
	maxUpdater []int64
}

func NewBucket(windowStart int64) *Bucket {

	adder := make([]int64, EventCount)
	maxUpdater := make([]int64, EventCount)
	return &Bucket{
		windowStart: windowStart,
		adder:       adder,
		maxUpdater:  maxUpdater,
	}
}

func (b *Bucket) GetAdder(event Event) int64 {

	return b.adder[event]

}

func (b *Bucket) GetMaxUpdater(event Event) int64 {

	return b.maxUpdater[event]

}
