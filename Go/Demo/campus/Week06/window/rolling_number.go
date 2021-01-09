package window

import (
	"fmt"
	"sync"
	"time"
)

type RollingNumber struct {
	timeInMilliseconds      int
	numberOfBuckets         int
	bucketSizeInMillseconds int
	buckets                 *CycleQueue
	mux                     sync.RWMutex
}

func NewRollingNumber(timeInMilliseconds, numberOfBuckets int) *RollingNumber {

	if timeInMilliseconds%numberOfBuckets != 0 {
		fmt.Println("timeInMilliseconds 必须能被 numberOfBuckets 整除. 例如 1000/10 可以, 1000/11 is 不行.")
		return nil
	}

	return &RollingNumber{
		timeInMilliseconds:      timeInMilliseconds,
		numberOfBuckets:         numberOfBuckets,
		bucketSizeInMillseconds: timeInMilliseconds / numberOfBuckets,
		buckets:                 NewCycleQueue(numberOfBuckets),
	}
}

/// 当前Bucket 自加1
func (r *RollingNumber) Increment(event Event) {
	r.GetCurrentBucket().Increment(event)
}

/// 当前Bucket 加上指定值
func (r *RollingNumber) Add(event Event, value int32) {
	r.GetCurrentBucket().Add(event, value)
}

// 更新当前maxUpdater，保留最大值
func (r *RollingNumber) UpdateRollingMax(event Event, value int32) {
	r.GetCurrentBucket().UpdateMaxUpdater(event, value)
}

// 清空数据
func (r *RollingNumber) Reset() {
	// 清空环形队列
	r.buckets.Clear()
}

//根据event type 获取所有Bucket 某index 总和
func (r *RollingNumber) GetRollingSum(event Event) int32 {
	if r.GetCurrentBucket() == nil {
		return 0
	}

	var sum int32 = 0
	for _, b := range r.buckets.data {
		bucket := b.(*Bucket)
		sum += bucket.GetAdder(event)
	}
	return sum

}

// 获取最后一个bucket 值
func (r *RollingNumber) GetValueOfLatestBucket(event Event) int32 {

	return r.buckets.GetLast().(*Bucket).GetAdder(event)
}

// 获取所有bucket 某一个索引的所有值
func (r *RollingNumber) GetValues(event Event) []int32 {

	result := make([]int32, r.buckets.curSize())
	for idx, b := range r.buckets.data {
		bucket := b.(*Bucket)
		result[idx] += bucket.GetAdder(event)
	}
	return result
}

// getValues 结果的最大值
func (r *RollingNumber) GetRollingMaxValue(event Event) int32 {

	result := r.GetValues(event)
	r.BubbleSort(result)
	return result[len(result)-1]
}

// 获取当前bucket
func (r *RollingNumber) GetCurrentBucket() *Bucket {
	currentTime := time.Now().Unix()

	var bucket *Bucket = r.buckets.GetLast().(*Bucket)
	if bucket != nil && currentTime < bucket.windowStart+(int64(r.bucketSizeInMillseconds)) {
		return bucket
	}

	// 如果为空，重新生成一个
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.buckets.GetLast() == nil {

		bucket = NewBucket(currentTime)
		r.buckets.Push(bucket)
		return bucket

	}

	for i := 0; i < r.buckets.curSize(); i++ {
		bucket = r.buckets.GetLast().(*Bucket)
		if currentTime < bucket.windowStart+(int64(r.bucketSizeInMillseconds)) {
			// 在窗口时间内，返回 bucket
			return bucket
		} else if currentTime-bucket.windowStart+(int64(r.bucketSizeInMillseconds)) > int64(r.timeInMilliseconds) {
			// 当前时间超过窗口范围，重置，重新获取
			r.Reset()
			r.GetCurrentBucket()
		} else {
			bucket = NewBucket(currentTime)
			r.buckets.Push(bucket)
		}

	}

	return r.buckets.GetLast().(*Bucket)
}

//简单的冒泡排序
func (r *RollingNumber) BubbleSort(nums []int32) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
