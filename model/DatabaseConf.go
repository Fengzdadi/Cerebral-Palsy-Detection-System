package model

type User struct {
	Userid   int    `bson:"Userid"`
	Username string `bson:"Username"`
	Age      int    `bson:"Age"`
	Gender   int    `bson:"Gender"`
	Phone    string `bson:"Phone"`
	Email    string `bson:"Email"`
	Password string `bson:"Password"`
}

type Result struct {
	VideoName   int    `bson:"VideoName"`
	Userid      int    `bson:"Userid"`
	VideoPath   string `bson:"VideoPath"`
	VideoRes    string `bson:"VideoRes"`
	Probability string `bson:"Probability"`
}
