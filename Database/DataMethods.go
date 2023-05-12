package Database

import (
	"Cerebral-Palsy-Detection-System/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func UserCheck(UserID string, UserPassword string) int {
	// 取出连接池中的连接
	conn, err := Pool.Get()
	defer Pool.Release(conn)
	if conn == nil {
		log.Fatal(err)
		return 0
	}
	if err != nil {
		log.Fatal(err)
		return 0
	}
	// 取出用户密码
	// 从数据库中取出用户密码
	var result model.User
	collection := conn.Database("CPDS_TEST").Collection("User")
	if collection == nil {
		log.Fatal(err)
		return 0
	}
	err = collection.FindOne(context.Background(), bson.M{"Userid": UserID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0
		}
		return 0
	}
	if result.Password == UserPassword {
		return 1
	} else {
		return 0
	}
}

func GetUserinfo(name string, user *model.User) {
	conn, err := Pool.Get()
	defer Pool.Release(conn)
	if conn == nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	collection := conn.Database("CPDS_TEST").Collection("User")
	if collection == nil {
		log.Fatal(err)
	}

	filter := bson.D{{"Username", name}}

	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	fmt.Print(user)
	return
}
