package rollingNumber

type rollingNumber struct {
	timeInMilliseconds int
	numberOfBuckets    int
	buckets            *CycleQueue
}

func NewRollingNumber(timeInMilliseconds, numberOfBuckets int) *rollingNumber {
	return &rollingNumber{
		timeInMilliseconds: timeInMilliseconds,
		numberOfBuckets:    numberOfBuckets,
		buckets:            NewCycleQueue(numberOfBuckets),
	}
}

/// 当前Bucket 自加1
func (r *rollingNumber) increment(event Event) {

}

/// 当前Bucket 加上指定值
func (r *rollingNumber) add(event Event, value int64) {

}

/**
 * @Description: 更新当前maxUpdater，保留最大值
 * @param event
 * @param value
 */
func (r *rollingNumber) updateRollingMax(event Event, value int64) {

}

func (r *rollingNumber) reset() {

}

//根据event type 获取所有Bucket 某index 总和
func (r *rollingNumber) getRollingSum(event Event) int64 {

	return 0
}

// 获取最后一个bucket 值
func (r *rollingNumber) getValueOfLatestBucket(event Event) int64 {
	return 0
}

// 获取所有bucket 某一个索引的所有值
func (r *rollingNumber) getValues(event Event) []int64 {

	return nil
}

// getValues 结果的最大值
func (r *rollingNumber) getRollingMaxValue(event Event) int64 {

	return 0
}

func (r *rollingNumber) getCurrentBucket() *Bucket {

	return nil
}
