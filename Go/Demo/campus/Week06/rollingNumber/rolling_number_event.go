package rollingNumber

type RollingNumberEvent int8

// 类型总数
const RollingNumberCount = 4

const (
	EVENT_SUCCESS      RollingNumberEvent = 0
	EVENT_FAILURE      RollingNumberEvent = 1
	EVENT_TIMEOUT      RollingNumberEvent = 2
	EVENT_REJECTION    RollingNumberEvent = 3
)

