package BaseInfoHis

import (
	"Cerebral-Palsy-Detection-System/WS/model"
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

func (b *BaseInfoHis) GetBaseInfoHis(belongToChildID int) ([]BaseInfoHis, error) {
	var bInfoHis []BaseInfoHis
	err := model.DB.Table("BaseInfoHis").Where("BelongToChildID = ?", belongToChildID).Find(&bInfoHis).Error
	return bInfoHis, err
}

func (b *BaseInfoHis) AddBaseInfoHis() error {
	err := model.DB.Table("BaseInfoHis").Create(&b).Error
	return err
}
