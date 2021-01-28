package model

type WsProtocol struct {
	CommandType string `json:"commandType"` //命令类型,进入路由的
	Body        string `json:"body"`
	Time        int64  `json:"time"`
}
