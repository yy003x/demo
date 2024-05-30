package po

import (
	"time"
)

type DmActivity struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id;comment:自增id"`
	ActID     string    `gorm:"unique;not null;default:'';size:32;column:act_id;comment:活动ID"`
	ActName   string    `gorm:"not null;default:'';size:30;column:act_name;comment:活动名称"`
	ActType   int       `gorm:"not null;default:0;column:act_type;comment:活动类型"`
	ActStatus int       `gorm:"not null;default:0;column:act_status;comment:活动状态"`
	StartTime time.Time `gorm:"not null;default:'1970-01-01 00:00:01';column:start_time;comment:开始时间"`
	EndTime   time.Time `gorm:"not null;default:'1970-01-01 00:00:01';column:end_time;comment:结束时间"`
	Version   int       `gorm:"default:0;column:version;comment:版本号"`
	CreatedAt time.Time `gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:更新时间"`
}

func (d DmActivity) TableName() string {
	return "dm_activity"
}

func (d DmActivity) GetID() int64 {
	return d.ID
}
