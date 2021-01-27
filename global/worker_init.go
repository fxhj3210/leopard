package global

import (
	"fmt"
	"leopard/websocket"
)

func workerWsInit() *websocket.Server {
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)
	Log.Info("WorkerWs初始化成功")
	return ws
}
func handleConnection(c websocket.Connection) {
	fmt.Println("client connected,id=", c.ID())
	c.Write(1, []byte("welcome client"))
	// 从浏览器中读取事件
	c.On("chat", func(msg string) {
		// 将消息打印到控制台，c .Context（）是iris的http上下文。
		fmt.Printf("%s sent: %s\n", c.Context().ClientIP(), msg)
		// 将消息写回客户端消息所有者：
		// c.Emit("chat", msg)
		c.To(websocket.All).Emit("chat", msg)
	})

	c.OnMessage(func(msg []byte) {
		fmt.Println("received msg:", string(msg))
		c.Write(1, []byte("hello aa"))
		c.To(websocket.All).Emit("chat", msg)
	})

	c.OnDisconnect(func() {
		fmt.Println("client Disconnect,id=", c.ID())
	})
}
