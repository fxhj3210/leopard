package model

var WorkerWsList map[string]Worker //前面的string便是当前的ws连接号

type WsProtocol struct {
	CommandType string `json:"commandType"` //命令类型,进入路由的
	Body        string `json:"body"`
	Time        int64  `json:"time"`
}
type Worker struct {
	Info WorkerInfo `json:"workerInfo"`
}

type WorkerInfo struct {
	Nickname   string `json:"nickname"`
	InternetIp string `json:"internetIp"`
	LanIp      string `json:"lanIp"`
	Free       bool   `json:"free"`
	CpuLoad    int64  `json:"cpuLoad"`
	MemoryLoad int64  `json:"memory"`
}
