package rollingNumber

/// 环形队列
type CycleQueue struct {
	data    []interface{} //存储空间
	head    int           //前指针,前指针负责弹出数据移动
	tail    int           //尾指针,后指针负责添加数据移动
	maxSize int           //设置切片最大容量
}

func NewCycleQueue(maxSize int) *CycleQueue {
	return &CycleQueue{
		data:    make([]interface{}, maxSize),
		maxSize: maxSize,
		head:    0,
		tail:    0,
	}
}

//入队操作
//判断队列是否队满,队满则新加的元素覆盖 head
func (q *CycleQueue) Push(data interface{}) {

	//将元素放入队列尾部
	q.data[q.tail] = data
	//尾部元素指向下一个空间位置
	q.tail = (q.tail + 1) % q.maxSize

	if q.IsFull() {
		//队列满时覆盖head
		q.head = q.tail + 1
	}

}

// 出队操作
func (q *CycleQueue) Pop() interface{} {
	if q.tail == q.head {
		return nil
	}
	data := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) % q.maxSize
	return data

}

// 队列是否满了
func (q *CycleQueue) IsFull() bool {
	return (q.tail+1)%q.maxSize == q.head
}

// 队列是否空了
func (q *CycleQueue) IsEmpty() bool {
	return q.tail == q.head
}

// 当前数量
func (q *CycleQueue) CurSize() int {
	return (q.tail - q.head + q.maxSize) % q.maxSize
}

// 剩余数量
func (q *CycleQueue) RemainSize() int {
	return q.maxSize - 1 - q.CurSize()
}
