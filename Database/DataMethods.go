package Database

import (
	"Cerebral-Palsy-Detection-System/model"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func UserCheck(UserID string, UserPassword string) bool {
	// 取出连接池中的连接
	conn, err := Pool.Get()
	defer Pool.Release(conn)
	if conn == nil {
		log.Fatal(err)
		return false
	}
	if err != nil {
		log.Fatal(err)
		return false
	}
	// 取出用户密码
	// 从数据库中取出用户密码
	var result model.User
	collection := conn.Database("CPDS").Collection("User")
	if collection == nil {
		log.Fatal(err)
		return false
	}
	err = collection.FindOne(context.Background(), bson.M{"Userid": UserID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if result.Password == UserPassword {
		return true
	} else {
		return false
	}
}

func GetUserinfo(name string) []byte {
	conn, err := Pool.Get()
	defer Pool.Release(conn)
	if conn == nil {
		log.Fatal(err)
		return nil
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}

	collection := conn.Database("CPDS").Collection("User")
	if collection == nil {
		log.Fatal(err)
		return nil
	}

	filter := bson.D{{"name", name}}

	var user model.User

	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return userJson
}
