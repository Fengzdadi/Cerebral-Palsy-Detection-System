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
	//Pool, err = NewConnPool(20)
	if err != nil {
		log.Fatal(err)
	}
}

// NewConnPool 初始化mongoDB连接池
func NewConnPool(maxConn int) (*ConnPool, error) {
	conns := make(chan *mongo.Client, maxConn)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for i := 0; i < maxConn; i++ {
		conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
		// databases, _ := conn.ListDatabaseNames(ctx, bson.M{})
		// log.Printf("%v", databases)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		conns <- conn
	}
	return &ConnPool{conns: conns}, nil
}

// Get 从连接池中获取一个连接
func (p *ConnPool) Get() (*mongo.Client, error) {
	select {
	case db := <-p.conns:
		return db, nil
	default:
		return nil, nil
	}
}

// Release 将连接放回连接池中
func (p *ConnPool) Release(db *mongo.Client) {
	p.conns <- db
}
