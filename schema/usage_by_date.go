package schema

type UsageByDate struct {
	RuUsed               int64 `gorm:"column:ru_used"`
	APICalls             int64 `gorm:"column:api_calls"`
	ConsumptionTimestamp int64 `gorm:"column:consumption_timestamp"`
}
