package model

import (
	"echo/module/log"
	. "echo/module/orm"
	"fmt"
)

type Banner struct {
	Id   int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name string `gorm:"column:name"`
}

func (b *Banner) GetBannerList() *Banner {
	var banners Banner

	DB().First(&banners, 1)
	fmt.Println(banners)

	if err := DB().First(&banners, 1).Error; err != nil {
		log.Debugf("Get banner error: %v", err)
		return &banners
	}

	return &banners
}

// set User's table name to be `profiles`
func (b Banner) TableName() string {
	return "banner"
}