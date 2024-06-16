package service

type Trainer struct {
	Content   string `bson:"content"`   // 内容
	StartTime int64  `bson:"startTime"` // 创建时间
	EndTime   int64  `bson:"endTime"`   // 过期时间
	Read      bool   `bson:"read"`      // 已读
}

type Result struct {
	StartTime int64
	Msg       string
	Content   interface{}
	From      string
}
