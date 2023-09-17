package model

import (
	"Cerebral-Palsy-Detection-System/Pkg/e"
	"Cerebral-Palsy-Detection-System/Serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"time"
)

type BaseInfoHis struct {
	gorm.Model
	BelongToChildID uint      `gorm:"BelongToChildID" form:"belongToChildID"`
	Time            time.Time `gorm:"Time" form:"time"`
	Height          float64   `gorm:"Height" form:"height"`
	Weight          float64   `gorm:"Weight" form:"weight"`
}

func GetBaseInfoHis(belongToChildID uint) ([]BaseInfoHis, Serializer.Response) {
	var bInfoHis []BaseInfoHis
	code := e.SUCCESS
	err := DB.Table("base_info_his").Where("belong_to_child_id = ?", belongToChildID).Find(&bInfoHis).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR
		return bInfoHis, Serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: "获取失败",
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
	err := DB.Table("base_info_his").Create(&b).Error
	if err != nil {
		code = e.ERROR
		return Serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: "创建失败",
		}
	} else {
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
}
