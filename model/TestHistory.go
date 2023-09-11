// Test x

package model

import (
	"Cerebral-Palsy-Detection-System/Pkg/e"
	"Cerebral-Palsy-Detection-System/Serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"time"
)

type TestHistory struct {
	gorm            gorm.Model
	BelongToChildID int       `gorm:"BelongToChildID"`
	TestTime        time.Time `gorm:"TestDate"`
	RawPath         string    `gorm:"RawPath"`
	ResPath         string    `gorm:"ResPath"`
	ResProbability  float64   `gorm:"ResProbability"`
}

func GetTestHistory(belongToChildID int) ([]TestHistory, Serializer.Response) {
	var tHistory []TestHistory
	code := e.SUCCESS
	err := DB.Table("TestHistory").Where("belong_to_children_id = ?", belongToChildID).Find(&tHistory).Error
	if err != nil {
		logging.Info(err)
		code := e.ERROR
		return tHistory, Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	} else {
		return tHistory, Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
}

// AddTestHistory function need to fix
func (t *TestHistory) AddTestHistory() Serializer.Response {
	code := e.SUCCESS
	err := DB.Table("TestHistory").Create(&t).Error
	if err != nil {
		code = e.ERROR
		return Serializer.Response{
			Code: code, // code fix
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
