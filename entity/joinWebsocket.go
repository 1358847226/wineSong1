package entity

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type JoinWebsocket struct {
	Token       string      `json:"token"`
	RoomId      int         `json:"roomId"`
	LeaveRoomId int         `json:"leaveRoomId"`
	Out         string      `json:"out"`
	Peng        string      `json:"peng"`
	Eat         string      `json:"eat"`
	Gang        string      `json:"gang"`
	Hu          string      `json:"hu"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

//客户端管理
type ClientManager struct {
	//客户端 map 储存并管理所有的长连接client，在线的为true，不在的为false
	Clients map[*Client]bool
	//web端发送来的的message我们用broadcast来接收，并最后分发给所有的client
	Broadcast chan []byte
	//谁发来的信息
	Cid string
	//新创建的长连接client
	Register chan *Client
	//新注销的长连接client
	Unregister chan *Client
}

//客户端 Client
type Client struct {
	//用户id
	Id string
	//连接的socket
	Socket *websocket.Conn
	//发送的消息
	Send chan []byte
}

//会把Message格式化成json
type Message struct {
	//消息struct
	Sender    string `json:"sender"`    //发送者
	Recipient string `json:"recipient"` //接收者
	Content   string `json:"content"`   //内容
}

//定义客户端管理的send方法
func (manager *ClientManager) Send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		//不给屏蔽的连接发送消息
		if conn != ignore {
			conn.Send <- message
		}
	}
}

func (manager *ClientManager) Start() {
	fmt.Println("start")
	for {
		select {
		//如果有新的连接接入,就通过channel把连接传递给conn
		case conn := <-manager.Register:
			//把客户端的连接设置为true
			manager.Clients[conn] = true
			//把返回连接成功的消息json格式化
			jsonMessage, _ := json.Marshal(&Message{Content: "id:" + manager.Cid + " has connected."})
			//调用客户端的send方法，发送消息
			manager.Send(jsonMessage, conn)
			//如果连接断开了
		case conn := <-manager.Unregister:
			//判断连接的状态，如果是true,就关闭send，删除连接client的值
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "id:" + manager.Cid + " has disconnected."})
				manager.Send(jsonMessage, conn)
			}
			//广播
		case message := <-manager.Broadcast:
			//遍历已经连接的客户端，把消息发送给他们
			//fmt.Println(message)
			for conn := range manager.Clients {
				//判断发送给谁（不发送自己）
				//接收者id： conn.id
				//发送者id： manager.cid
				if conn.Id == manager.Cid {
					continue
				}
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
	Cid:        "",
}
var (
	Upgrader = websocket.Upgrader{
		//
		ReadBufferSize: 1023,
		//
		WriteBufferSize: 1023,
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//定义客户端结构体的read方法
func (c *Client) Read() {
	defer func() {
		//结构体cid赋值
		Manager.Cid = c.Id
		//触发关闭
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		//读取消息
		_, message, err := c.Socket.ReadMessage()
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		//如果没有错误信息就把信息放入broadcast
		jsonMessage, _ := json.Marshal(&Message{Sender: c.Id, Content: string(message)})
		//结构体cid赋值
		Manager.Cid = c.Id
		//触发消息发送
		Manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		//从send里读消息
		case message, ok := <-c.Send:
			//如果没有消息
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//有消息就写入，发送给web端
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
