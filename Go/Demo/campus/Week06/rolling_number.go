package main

import (
	"sort"
	"time"
)

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
	r.getCurrentBucket().increment(event)
}

/// 当前Bucket 加上指定值
func (r *rollingNumber) add(event Event, value int) {
	r.getCurrentBucket().add(event, value)
}

/**
 * @Description: 更新当前maxUpdater，保留最大值
 * @param event
 * @param value
 */
func (r *rollingNumber) updateRollingMax(event Event, value int) {
	r.getCurrentBucket().updateMaxUpdater(event, value)
}

func (r *rollingNumber) reset() {

}

//根据event type 获取所有Bucket 某index 总和
func (r *rollingNumber) getRollingSum(event Event) int {
	if r.getCurrentBucket() == nil {
		return 0
	}

	sum := 0
	for _, b := range r.buckets.data {
		bucket := b.(*Bucket)
		sum += bucket.GetAdder(event)
	}
	return sum

}

// 获取最后一个bucket 值
func (r *rollingNumber) getValueOfLatestBucket(event Event) int {

	return r.buckets.getLast().(*Bucket).GetAdder(event)
}

// 获取所有bucket 某一个索引的所有值
func (r *rollingNumber) getValues(event Event) []int {

	result := make([]int, r.buckets.curSize())
	for idx, b := range r.buckets.data {
		bucket := b.(*Bucket)
		result[idx] += bucket.GetAdder(event)
	}
	return result
}

// getValues 结果的最大值
func (r *rollingNumber) getRollingMaxValue(event Event) int {

	result := r.getValues(event)
	sort.Ints(result)
	return result[len(result)-1]
}

func (r *rollingNumber) getCurrentBucket() *Bucket {
	var bucket *Bucket
	if r.buckets.getLast() == nil {
		// 如果为空，重新生成一个
		currentTime := time.Now().Second()
		bucket = NewBucket(currentTime)
	} else {
		bucket = r.buckets.getLast().(*Bucket)
	}
	return bucket
}
