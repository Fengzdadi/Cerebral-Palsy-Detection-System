package model

import (
	"Cerebral-Palsy-Detection-System/WS/Pkg/e"
	"Cerebral-Palsy-Detection-System/WS/Serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"time"
)

type BaseInfoHis struct {
	gorm            gorm.Model
	BelongToChildID int       `gorm:"BelongToChildID"`
	Time            time.Time `gorm:"Time"`
	height          float64   `gorm:"height"`
	weight          float64   `gorm:"weight"`
}

func GetBaseInfoHis(belongToChildID int) ([]BaseInfoHis, Serializer.Response) {
	var bInfoHis []BaseInfoHis
	code := e.SUCCESS
	err := DB.Table("BaseInfoHis").Where("BelongToChildID = ?", belongToChildID).Find(&bInfoHis).Error
	if err != nil {
		logging.Info(err)
		code := e.ERROR
		return bInfoHis, Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	} else {
		return bInfoHis, Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
}

func (b *BaseInfoHis) AddBaseInfoHis() Serializer.Response {
	code := e.SUCCESS
	err := DB.Table("BaseInfoHis").Create(&b).Error
	if err != nil {
		code = e.ERROR
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "",
		}
	} else {
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
}
