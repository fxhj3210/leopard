package api

import (
	"leopard/model"
	"leopard/service"
	"leopard/websocket"
)

// Identifying 表明用户身份,首次注册
func Identifying(P model.WsProtocol, Ws websocket.Connection) (err error) {
	var service service.WsWorkerInfo

	err = service.WsWorkerNewInfo(P, Ws)
	if err != nil {
		return err
	}
	return nil

}

// UpdateState 更新任务状态
func UpdateState(P model.WsProtocol, Ws websocket.Connection) (err error) {
	var service service.WsWorkerInfo

	err = service.WsWorkerUpdateState(P, Ws)
	if err != nil {
		return err
	}
	return nil
}
