package model

import (
	"Cerebral-Palsy-Detection-System/WS/Pkg/e"
	"Cerebral-Palsy-Detection-System/WS/Serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type ChildrenInfo struct {
	gorm.Model
	BelongTo  uint       `form:"BelongTo" bson:"BelongTo"`
	ChildName string     `form:"ChildName" bson:"ChildName"`
	Age       int        `form:"Age" bson:"Age"`
	Gender    GenderType `form:"Gender" bson:"Gender"`
	BirthDate string     `form:"BirthDate" bson:"BirthDate"`
}

type GenderType string

const (
	Male        GenderType = "Male"
	Female      GenderType = "Female"
	OtherGender GenderType = "OtherGender"
)

func GetChildInfo(belongTo uint) ([]ChildrenInfo, Serializer.Response) {
	var cInfo []ChildrenInfo
	code := e.SUCCESS
	err := DB.Table("children_info").Where("belong_to = ?", belongTo).Find(&cInfo).Error
	if err != nil {
		logging.Info(err)
		code := e.ERROR
		return cInfo, Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	} else {
		return cInfo, Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
}

func (c *ChildrenInfo) AddChildInfo(userid uint) Serializer.Response {
	// 判断是否已经存在
	var count int
	code := e.SUCCESS
	DB.Model(&ChildrenInfo{}).Where("child_name = ?", c.ChildName).Count(&count)
	// fmt.Println(count)
	if count > 0 {
		code := e.ERROR
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "已经存在该用户了",
		}
	}

	rand.Seed(time.Now().UnixNano())
	c.ID = uint(rand.Intn(100000000))

	c.BelongTo = userid

	if err := DB.Table("children_info").Create(&c).Error; err != nil {
		logging.Info(err)
		code := e.ErrorDatabase
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return Serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
	}
}
