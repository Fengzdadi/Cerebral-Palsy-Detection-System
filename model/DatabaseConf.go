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
