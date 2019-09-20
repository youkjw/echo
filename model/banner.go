package model

import (
	"time"
	. "echo/module/orm"
	"echo/module/log"
)

func (b *Banner) GetBannerList() *Banner {
	banner := Banner{}

	if err := DB().Take(&banner).Error; err != nil {
		log.Debugf("Get banner error: %v", err)
		return &banner
	}

	return &banner
}

type Banner struct {
	Id          uint64    `gorm:"AUTO_INCREMENT"`
	Banner      string    `gorm:"column:banner"`
	Link        string    `gorm:"column:link"`
	Status      uint64    `gorm:"column:status"`
	CreatedTime uint64    `gorm:"column:created_time"`
	Type        time.Time `gorm:"column:type"`
	Desc        time.Time `gorm:"column:desc"`
	Sort        time.Time `gorm:"column:sort"`
	Popup       time.Time `gorm:"column:popup"`
}

func (b Banner) TableName() string {
	return "yi_banner"
}