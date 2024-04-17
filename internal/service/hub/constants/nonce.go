package constants

import "time"

const (
	NonceLife      = 5 * time.Minute
	NonceKeyPrefix = "paymentprocessor:nonce"
)
