package model

/*
WorkerInfo
这个仅存在于内存中,不涉及数据库,当建立WebSocket连接时,首次更新信息
也就是说等待客户端主动连接,客户端自动寻找同网段的服务器.客户端一旦与服务器取得连接,自动汇报自己的情况
*/
type WorkerInfo struct {
	ID   string
	free bool //空闲状态

}
