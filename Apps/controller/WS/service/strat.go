package service

import (
	"Cerebral-Palsy-Detection-System/Apps/controller/WS/Conf"
	"Cerebral-Palsy-Detection-System/Pkg/e"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var lock sync.RWMutex

func (manager *ClientManager) Start() {
	log.Println("<---监听管道通信--->")
	for {
		select {
		case c := <-Manager.Register: // 建立连接
			log.Printf("建立新连接: %v", c.ID)
			lock.Lock()
			Manager.Clients[c.ID] = c
			lock.Unlock()
			replyMsg := &ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: "已连接至服务器",
			}
			// 设置之前消息全部read
			SetOneRead(Conf.MongoDBName, c.SendID)
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister: // 断开连接
			log.Printf("断开连接:%v", conn.ID)
			lock.Lock()
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketEnd,
					Content: "连接已断开",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
			lock.Unlock()
		case c := <-Manager.Reply:
			if s, ok := Manager.Clients[c.To]; !ok {
				continue
			} else {
				// 可能需要设置已读？
				s.Send <- []byte(c.Content)
			}
		}
	}
}
