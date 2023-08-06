package TestHistory

import (
	"Cerebral-Palsy-Detection-System/WS/model"
	"github.com/jinzhu/gorm"
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

func (t *TestHistory) GetTestHistory(belongToChildID int) ([]TestHistory, error) {
	var tHistory []TestHistory
	err := model.DB.Table("TestHistory").Where("BelongToChildID = ?", belongToChildID).Find(&tHistory).Error
	return tHistory, err
}

func (t *TestHistory) AddTestHistory() error {
	err := model.DB.Table("TestHistory").Create(&t).Error
	return err
}
