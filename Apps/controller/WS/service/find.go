package service

import (
	"Cerebral-Palsy-Detection-System/Apps/controller/WS/Conf"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sort"
	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     bool   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

func InsertMsg(database string, id string, content string, expire int64) (err error) {
	collection := Conf.MongoDBClient.Database(database).Collection(id)
	comment := Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      false,
	}
	_, err = collection.InsertOne(context.TODO(), &comment)
	return
}

// 每次用户连接到websocket，将对方发送的消息置为已读
func SetOneRead(database string, sendID string) (err error) {
	collection := Conf.MongoDBClient.Database(database).Collection(sendID)
	filter := bson.M{"read": false}

	// 设置要更新的字段
	update := bson.M{"$set": bson.M{"read": true}}

	// 执行更新操作
	updateResult, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("更新了 %d 个文档\n", updateResult.ModifiedCount)
	return nil
}

func FindMany(database string, sendID string, id string, time int64, pageSize int) (results []Result, err error) {
	var resultsMe []Trainer
	var resultsYou []Trainer
	sendIDCollection := Conf.MongoDBClient.Database(database).Collection(sendID)
	idCollection := Conf.MongoDBClient.Database(database).Collection(id)

	//cur, err := sendIDCollection.Find(context.TODO(), bson.D{})

	sendIDTimeCursor, err := sendIDCollection.Find(context.TODO(), bson.D{},
		options.Find().SetSort(bson.D{{"startTime", -1}}).SetLimit(int64(pageSize)))

	idTimeCursor, err := idCollection.Find(context.TODO(), bson.D{},
		options.Find().SetSort(bson.D{{"startTime", -1}}).SetLimit(int64(pageSize)))

	err = sendIDTimeCursor.All(context.TODO(), &resultsYou)
	err = idTimeCursor.All(context.TODO(), &resultsMe)
	fmt.Println(resultsYou)
	fmt.Println(resultsMe)
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return results, nil
}

//func FirstFindtMsg(database string, sendId string, id string) (results []Result, err error) {
//	// 首次查询(把对方发来的所有未读都取出来)
//	var resultsMe []Trainer
//	var resultsYou []Trainer
//	sendIdCollection := Conf.MongoDBClient.Database(database).Collection(sendId)
//	idCollection := Conf.MongoDBClient.Database(database).Collection(id)
//	filter := bson.M{"read": bson.M{
//		"&all": []uint{0},
//	}}
//	sendIdCursor, err := sendIdCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{
//		"startTime", 1}}), options.Find().SetLimit(1))
//	if sendIdCursor == nil {
//		return
//	}
//	var unReads []Trainer
//	err = sendIdCursor.All(context.TODO(), &unReads)
//	if err != nil {
//		log.Println("sendIdCursor err", err)
//	}
//	if len(unReads) > 0 {
//		timeFilter := bson.M{
//			"startTime": bson.M{
//				"$gte": unReads[0].StartTime,
//			},
//		}
//		sendIdTimeCursor, _ := sendIdCollection.Find(context.TODO(), timeFilter)
//		idTimeCursor, _ := idCollection.Find(context.TODO(), timeFilter)
//		err = sendIdTimeCursor.All(context.TODO(), &resultsYou)
//		err = idTimeCursor.All(context.TODO(), &resultsMe)
//		results, err = AppendAndSort(resultsMe, resultsYou)
//	} else {
//		results, err = FindMany(database, sendId, id, 9999999999, 10)
//	}
//	overTimeFilter := bson.D{
//		{"$and", bson.A{
//			bson.D{{"endTime", bson.M{"&lt": time.Now().Unix()}}},
//			bson.D{{"read", bson.M{"$eq": 1}}},
//		}},
//	}
//	_, _ = sendIdCollection.DeleteMany(context.TODO(), overTimeFilter)
//	_, _ = idCollection.DeleteMany(context.TODO(), overTimeFilter)
//	// 将所有的维度设置为已读
//	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
//		"$set": bson.M{"read": 1},
//	})
//	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
//		"&set": bson.M{"ebdTime": time.Now().Unix() + int64(3*month)},
//	})
//	return
//}

func AppendAndSort(resultsMe, resultsYou []Trainer) (results []Result, err error) {
	for _, r := range resultsMe {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}
	for _, r := range resultsYou {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "you",
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}
