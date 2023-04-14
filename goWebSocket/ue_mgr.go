package goWebSocket

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type WebWsUser struct {
	ID     string
	Socket *wsConn
	Send   chan []byte
}

type UserManager struct {
	Users      map[*WebWsUser]bool
	Broadcase  chan []byte
	Register   chan *WebWsUser
	Unregister chan *WebWsUser
}

type wsMgr = UserManager
type wsFd = WebWsUser

var RecvMesChan chan []byte

var LogWsMgr wsMgr

func init_mes_chan() {
	if RecvMesChan == nil {
		RecvMesChan = make(chan []byte, 20)
	}
}

func (mgr *UserManager) RegClientUser(id string, conn *wsConn) {
	user := &WebWsUser{
		ID:     id,
		Socket: conn,
		Send:   make(chan []byte, 100),
	}
	//every connect user init a 100 channel send chan
	go user.Write()
	mgr.Register <- user
}

func (mgr *UserManager) Start(fire <-chan []byte) {
	for {
		select {
		case conn := <-mgr.Register:
			mgr.Users[conn] = true
			fmt.Println("The regClient num is ", len(mgr.Users))
			mgr.SendOnly([]byte("ok a client had connected"), conn)
		case user := <-mgr.Unregister:
			if _, ok := mgr.Users[user]; ok {
				mgr.Send([]byte("a client has disconnected"))
				close(user.Send)
				delete(mgr.Users, user)
			}
		case msg := <-fire:
			for conn := range mgr.Users {
				select {
				case conn.Send <- msg:
					fmt.Println("recv upload message", string(msg))
				default:
					fmt.Println("close Send and Delete Users connect")
					close(conn.Send)
					delete(mgr.Users, conn)
				}
			}
		case msg := <-mgr.Broadcase:
			for conn := range mgr.Users {
				select {
				case conn.Send <- msg:
				default:
					close(conn.Send)
					delete(mgr.Users, conn)
				}
			}
		}
	}
}

func (mgr *UserManager) SendOnly(msg []byte, user *WebWsUser) {
	for conn := range mgr.Users {
		if conn == user {
			conn.Send <- msg
			return
		}
	}
}

func (mgr *UserManager) Send(msg []byte) {
	mgr.Broadcase <- msg
}

func (c *WebWsUser) Write() {
	defer func() {
		fmt.Println("The websocket session is close")
		_ = c.Socket.Close()
	}()
	fmt.Println("The range Process start run")
	for msg := range c.Send {
		c.Socket.WriteMessage(websocket.TextMessage, msg)
		//fmt.Println("Send message Error is ", err)
	}
	fmt.Println("The range Process stop run")
}
