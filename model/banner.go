package model

import (
	"time"
	. "echo/module/orm"
)

func (b *Banner) GetBannerList() *Banner {
	banner := Banner{}

	if err := DB().Where("id = ?", id).First(&post).Error; err != nil {
		log.Debugf("Get post error: %v", err)
		return nil
	}

	if err := DB().Model(&post).Related(&post.User).Error; err != nil {
		log.Debugf("Post user related error: %v", err)
		return &post
	}

	return &post
}

type Banner struct {
	Id          uint64    `gorm:"AUTO_INCREMENT"`
	Banner      string    `gorm:"column:Banner"`
	Link        string    `gorm:"column:link"`
	Status      uint64    `gorm:"column:status"`
	CreatedTime uint64    `gorm:"column:created_time"`
	Type        time.Time `gorm:"column:type"`
	Desc        time.Time `gorm:"column:desc"`
	Sort        time.Time `gorm:"column:sort"`
	Popup       time.Time `gorm:"column:popup"`
}
