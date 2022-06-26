package models

import "time"

type Staff struct {
	ID         int64      `gorm:column:id;PRIMERY_KEY`
	Username   string     `gorm:column:username`
	Password   string     `gorm:column:password`
	Fullname   string     `gorm:column:fullname`
	Salary     int64      `gorm:column:salary`
	CreateTime *time.Time `gorm:column:create_time`
	UpdateTime *time.Time `gorm:column:update_time`
}
