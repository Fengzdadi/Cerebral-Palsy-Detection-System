package model

import (
	"github.com/jinzhu/gorm"
)

type Kinship struct {
	gorm                gorm.Model
	UserID              int              `gorm:"UserID"`
	TargetID            int              `gorm:"TargetID"`
	Relationship        RelationshipType `gorm:"Relationship"`
	InverseRelationship RelationshipType `gorm:"InverseRelationship"`
	IsShare             bool             `gorm:"IsShare"`
}

type RelationshipType string

const (
	Father        RelationshipType = "Father"
	Mother        RelationshipType = "Mother"
	Grandfather   RelationshipType = "Grandfather"
	Grandmother   RelationshipType = "Grandmother"
	OtherKinship  RelationshipType = "OtherKinship"
	husband       RelationshipType = "husband"
	wife          RelationshipType = "wife"
	son           RelationshipType = "son"
	daughter      RelationshipType = "daughter"
	grandson      RelationshipType = "grandson"
	granddaughter RelationshipType = "granddaughter"
)

func GetKinship(userID int) ([]Kinship, error) {
	var kInfo []Kinship
	err := DB.Table("Kinship").Where("UserID = ?", userID).Find(&kInfo).Error
	return kInfo, err
}

func (k *Kinship) AddKinship() error {
	err := DB.Table("Kinship").Create(&k).Error
	return err
}
