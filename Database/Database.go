package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type ConnPool struct {
	conns chan *mongo.Client
}

var Pool *ConnPool

func DatabaseInit() {
	var err error
	Pool, err = NewConnPool(20)
	if err != nil {
		log.Fatal(err)
	}
}

// NewConnPool 初始化mongoDB连接池
func NewConnPool(maxConn int) (*ConnPool, error) {
	conns := make(chan *mongo.Client, maxConn)
	for i := 0; i < maxConn; i++ {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			return nil, err
		}
		conns <- client
	}
	return &ConnPool{conns: conns}, nil
}
