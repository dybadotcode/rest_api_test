package models

import (
	"time"
)

//Rss ...
type Rss struct {
	ID           int64     `gorm:"type:bigserial; primaryKey"`
	AccessTime   time.Time `gorm:"type:timestamp with time zone"`
	ResponseTime time.Time `gorm:"type:timestamp with time zone"`
	UserAgent    string    `gorm:"type:varchar"`
	URL          string    `gorm:"type:text"`
	Content      string    `gorm:"type:text"`
}
