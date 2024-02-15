package models

import "time"

type Url struct {
	Id          uint   `gorm:"primaryKey" type:"autoIncrement"`
	Long_url    string `form:"long_url"`
	Short_url   string
	Is_used     bool
	Expire_date time.Time
}
