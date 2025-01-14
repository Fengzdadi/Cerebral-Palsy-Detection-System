package service

import (
	"Cerebral-Palsy-Detection-System/Apps/controller/WS/Cache"
	"Cerebral-Palsy-Detection-System/Apps/controller/WS/Conf"
	"Cerebral-Palsy-Detection-System/Pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

const month = 60 * 60 * 24 * 30

// SendMsg 发送信息
type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ReplyMsg 回复信息
type ReplyMsg struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast // 广播通道
	Reply      chan *ReplyMsg  // 回复通道
	Register   chan *Client    // 注册通道
	Unregister chan *Client    // 注销通道
}

type Message struct {
	Sender    string `json:"sender,omitempty"` // omitempty表示如果为空则忽略
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client),
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *ReplyMsg),
	Unregister: make(chan *Client),
}

func createId(uid, toUid string) string {
	return uid + "->" + toUid
}

// ws://localhost:8080/ws?uid=123&toUid=456
func WsHandler(c *gin.Context) {
	uid := c.Query("uid")     // 自己的id
	toUid := c.Query("toUid") // 对方的id
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// 创建一个用户实例
	client := &Client{
		ID:     createId(uid, toUid),
		SendID: createId(toUid, uid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	// 用户注册到用户管理上
	Manager.Register <- client
	go client.Read()
	go client.Write()
}

func (c *Client) Read() {
	defer func() { // 避免忘记关闭，所以要加上close
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		// _,msg,_:=c.Socket.ReadMessage()
		err := c.Socket.ReadJSON(sendMsg) // 读取json格式，如果不是json格式，会报错
		if err != nil {
			log.Println("数据格式不正确", err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}

		if sendMsg.Type == 1 {
			r1, _ := Cache.RedisClient.Get(c.ID).Result()
			r2, _ := Cache.RedisClient.Get(c.SendID).Result()
			if r1 >= "3" && r2 == "" { // 限制单聊
				replyMsg := ReplyMsg{
					Code:    e.WebsocketLimit,
					Content: "达到限制",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				_, _ = Cache.RedisClient.Expire(c.ID, time.Hour*24*30).Result() // 防止重复骚扰，未建立连接刷新过期时间一个月
				continue
			} else {
				Cache.RedisClient.Incr(c.ID)
				_, _ = Cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result() // 防止过快“分手”，建立连接三个月过期
			}
			log.Println(c.ID, "发送消息", sendMsg.Content)
			// 插入发送消息
			InsertMsg(Conf.MongoDBName, c.ID, sendMsg.Content, int64(time.Hour*24*30))
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccess,
				To:      c.SendID,
				Content: sendMsg.Content,
			}
			Manager.Reply <- &replyMsg
		} else if sendMsg.Type == 2 { //拉取历史消息
			timeT, err := strconv.Atoi(sendMsg.Content) // 传送来时间
			if err != nil {
				timeT = 999999999
			}
			results, _ := FindMany(Conf.MongoDBName, c.SendID, c.ID, int64(timeT), 10)
			//if len(results) > 10 {
			//	results = results[:10]
			//}
			fmt.Println(results)
			msg, _ := json.Marshal(results)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		} else { // 找到另一个连接并发送消息

		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Println(c.ID, "接受消息:", string(message))
			replyMsg := ReplyMsg{
				//From: c.ID,
				//To: c.SendID,
				Code:    e.WebsocketSuccessMessage,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
