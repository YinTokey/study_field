package rollingNumber

type Event int8

// 事件类型总数
const EventCount = 4

const (
	EVENT_SUCCESS   Event = 0
	EVENT_FAILURE   Event = 1
	EVENT_TIMEOUT   Event = 2
	EVENT_REJECTION Event = 3
)
