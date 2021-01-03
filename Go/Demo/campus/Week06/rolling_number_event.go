package main

type RollingNumberEvent int32

const (
	EVENT_SUCCESS      RollingNumberEvent = 0
	EVENT_FAILURE      RollingNumberEvent = 1
	EVENT_TIMEOUT      RollingNumberEvent = 2
	EVENT_REJECTION    RollingNumberEvent = 3
)
