package model

import (
	"time"
)

//type User struct {
//	gorm.Model
//	Userid   int    `bson:"Userid"` // `"grom:"unique;not null`
//	Username string `bson:"Username"`
//	Age      int    `bson:"Age"`
//	Gender   int    `bson:"Gender"`
//	Phone    string `bson:"Phone"`
//	Email    string `bson:"Email"`
//	Password string `bson:"Password"`
//}

type SessionData struct {
	Username string `json:"Username"`
	Age      int    `json:"Age"`
	Gender   int    `json:"Gender"`
	Email    string `json:"Email"`
	Phone    string `json:"Phone"`
}

type VideoResult struct {
	VideoName   int    `bson:"VideoName"`
	Userid      int    `bson:"Userid"`
	VideoPath   string `bson:"VideoPath"`
	VideoRes    string `bson:"VideoRes"`
	Probability string `bson:"Probability"`
}

type Result struct {
	VideoName int       `bson:"VideoName"`
	Time      time.Time `bson:"Time"`
	Result    float64   `bson:"Result"`
	ResultAdd string    `bson:"ResultAdd"`
}

type HisResult struct {
	Userid     int      `bson:"Userid"`
	ResultData []Result `bson:"ResultData"`
	Count      int      `bson:"Count"`
}
