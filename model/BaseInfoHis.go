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
	BelongToChildID uint      `form:"BelongToChildID" gorm:"BelongToChildID"`
	Time            time.Time `form:"Time" gorm:"Time"`
	Height          float64   `form:"height" gorm:"height"`
	Weight          float64   `form:"weight" gorm:"weight"`
}

func GetBaseInfoHis(belongToChildID uint) ([]BaseInfoHis, Serializer.Response) {
	var bInfoHis []BaseInfoHis
	code := e.SUCCESS
	err := DB.Table("base_info_his").Where("belong_to_child_id = ?", belongToChildID).Find(&bInfoHis).Error
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
	b.Time = time.Now()
	code := e.SUCCESS
	err := DB.Table("base_info_his").Create(&b).Error
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
