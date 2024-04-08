package constants

import "time"

const (
	// EpochMutexLockKey : Mutex key for epoch billing execution
	EpochMutexLockKey = "paymentprocessor.lock.settlement"

	// EpochMutexExpiration : Maximal time of every execution (auto release when last execution fail to prevent stuck)
	EpochMutexExpiration = 1 * time.Hour // TODO: discuss a more detailed value
)
