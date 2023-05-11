package model

type User struct {
	Userid   int    `bson:"Userid"`
	Username string `bson:"Username"`
	Age      int    `bson:"Age"`
	Gender   int    `bson:"Gender"`
	phone    string `bson:"phone"`
	email    string `bson:"email"`
	Password string `bson:"Password"`
}
