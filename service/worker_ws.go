package service

import (
	"errors"
	"fmt"
	"leopard/global"
	"leopard/model"
	"leopard/serializer"
	"leopard/websocket"
)

var WorkerWs = WorkerWsInit()

func WorkerWsInit() *websocket.Server {
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)
	global.Log.Info("WorkerWs初始化成功")
	return ws
}

//控制器
func handleConnection(c websocket.Connection) {
	fmt.Println("client connected,id=", c.ID())
	c.Write(1, []byte("welcome client"))
	// 从浏览器中读取事件
	c.On("chat", func(msg string) {
		// 将消息打印到控制台，c .Context（）是iris的http上下文。
		fmt.Printf("%s sent: %s\n", c.Context().ClientIP(), msg)
		// 将消息写回客户端消息所有者：
		// c.Emit("chat", msg)
		c.Emit("chat", msg)
		//c.To(websocket.All).Emit("chat", msg)
	})

	c.OnMessage(func(msg []byte) {
		fmt.Println("received msg:", string(msg))
		//反序列化
		d, err := serializer.WsDeserialization(msg)
		if err != nil {
			c.Write(1, serializer.WsSerializerReq("err", err.Error()))
			return
		}

		//进入路由
		err = WorkerWsRouter(d)
		if err != nil {
			c.Write(1, serializer.WsSerializerReq("err", err.Error()))
			return
		}

		//c.Write(1, []byte(msg))
		//c.To(websocket.All).Emit("chat", msg)
	})

	c.OnError(func(err error) {
		global.Log.Warn("WorkerWsErr:" + err.Error())
		fmt.Println("client OnError,err=", err.Error())
	})
	c.OnDisconnect(func() {
		global.Log.Info("WorkerWsDisconnect:" + c.ID())
		fmt.Println("client Disconnect,id=", c.ID())
	})
}

//ReportState		向控制机汇报自己的状态,包括是否空闲
func WorkerWsRouter(p model.WsProtocol) (e error) {
	switch p.CommandType {
	case "reportState":

		break
	case "":
		break

	default:
		return errors.New("无法理解命令含义")

	}
	return nil
}
