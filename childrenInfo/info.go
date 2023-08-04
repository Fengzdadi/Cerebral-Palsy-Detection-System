package childrenInfo

import (
	"Cerebral-Palsy-Detection-System/WS/model"
	"github.com/jinzhu/gorm"
)

type ChildrenInfo struct {
	gorm      gorm.Model
	BelongTo  int        `bson:"BelongTo"`
	ChildName string     `bson:"ChildName"`
	Age       int        `bson:"Age"`
	Gender    GenderType `bson:"Gender"`
	BirthDate string     `bson:"BirthDate"`
}

type GenderType string

const (
	Male   GenderType = "Male"
	Female GenderType = "Female"
	Other  GenderType = "Other"
)

func (c *ChildrenInfo) GetChildInfo(belongTo int) ([]ChildrenInfo, error) {
	var cInfo []ChildrenInfo
	err := model.DB.Table("ChildrenInfo").Where("BelongTo = ?", belongTo).Find(&cInfo).Error
	return cInfo, err
}

func (c *ChildrenInfo) AddChildInfo() error {
	err := model.DB.Table("ChildrenInfo").Create(&c).Error
	return err
}
