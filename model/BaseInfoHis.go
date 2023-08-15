package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BaseInfoHis struct {
	gorm            gorm.Model
	BelongToChildID int       `gorm:"BelongToChildID"`
	Time            time.Time `gorm:"Time"`
	height          float64   `gorm:"height"`
	weight          float64   `gorm:"weight"`
}

func GetBaseInfoHis(belongToChildID int) ([]BaseInfoHis, error) {
	var bInfoHis []BaseInfoHis
	err := DB.Table("BaseInfoHis").Where("BelongToChildID = ?", belongToChildID).Find(&bInfoHis).Error
	return bInfoHis, err
}

func (b *BaseInfoHis) AddBaseInfoHis() error {
	err := DB.Table("BaseInfoHis").Create(&b).Error
	return err
}
