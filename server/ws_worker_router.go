package server

import (
	"errors"
	"leopard/api"
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
	//让工作机表明身份
	global.Log.Warn("与Worker建立连接" + c.ID())
	c.Write(1, serializer.WsSerializerReq(model.WsProtocol{CommandType: "identifying"}))

	/*
		// 从浏览器中读取事件
		c.On("chat", func(msg string) {
			// 将消息打印到控制台，c .Context（）是iris的http上下文。
			fmt.Printf("%s sent: %s\n", c.Context().ClientIP(), msg)
			// 将消息写回客户端消息所有者：
			// c.Emit("chat", msg)
			c.Emit("chat", msg)
			//c.To(websocket.All).Emit("chat", msg)
		})
	*/

	c.OnMessage(func(msg []byte) {
		err := WorkerWsRouter(msg, c)
		if err != nil {
			c.Write(1, serializer.WsSerializerReq(model.WsProtocol{CommandType: "err", Body: err.Error()}))
			global.Log.Warn("WsID:"+c.ID()+"|", "Msg:"+string(msg)+"|", "Err:"+err.Error())
			return
		}
		c.Write(1, serializer.WsSerializerReq(model.WsProtocol{CommandType: "succeed"}))

		//c.Write(1, []byte(msg))
		//c.To(websocket.All).Emit("chat", msg)
	})

	c.OnError(func(err error) {
		global.Log.Warn("WorkerWsErr:" + err.Error())
	})
	c.OnDisconnect(func() {
		global.Log.Info("WorkerWsDisconnect:" + c.ID())
		cls(c)
	})
}

/*
===============
=====Ws路由=====
===============
	收
		-reportState		向控制机汇报自己的状态,包括是否空闲
		-identifying		表明连接进入的Ws客户端身份

	发
		-pushTask			向空闲的任务机推送任务
*/
//ReportState		向控制机汇报自己的状态,包括是否空闲
func WorkerWsRouter(msg []byte, Ws websocket.Connection) (err error) {

	//反序列化
	p, err := serializer.WsDeserialization(msg)
	if err != nil {
		return err
	}

	//进入路由
	switch p.CommandType {
	case "identifying":
		err = api.Identifying(*p, Ws)
		break
	case "UpdateState":
		err = api.UpdateState(*p, Ws)
		break
	default:
		return errors.New("无法理解命令含义")
	}
	if err != nil {
		return err
	}
	return nil
}

/*
====================
=====断开连接清理=====
====================
*/
func cls(Ws websocket.Connection) {
	global.Log.Info("与" + Ws.ID() + "断开Ws连接")
	delete(model.WorkerWsList, Ws.ID())
}
