package models

type Job struct {
	ID           string `gorm:"primaryKey"`
	Version      int    `gorm:"primaryKey"`
	UID          string
	Status       string
	Rate         float64
	Title        string
	CompanyID    string
	ContractorID string
}

type Timelog struct {
	ID             string `gorm:"primaryKey"`
	Version        int    `gorm:"primaryKey"`
	UID            string
	Duration       int
	TimestampStart int64
	TimestampEnd   int64
	Type           string
	JobUID         string
}

type PaymentLineItem struct {
	ID         string `gorm:"primaryKey"`
	Version    int    `gorm:"primaryKey"`
	UID        string
	JobUID     string
	TimelogUID string
	Amount     float64
	Status     string
}
