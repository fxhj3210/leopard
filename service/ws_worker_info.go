package service

import (
	"encoding/json"
	"errors"
	"leopard/global"
	"leopard/model"
	"leopard/websocket"
)

// WsWorkerInfo 更新用户信息
type WsWorkerInfo struct {
}

// WsWorkerNewInfo 初次使用,更新用户信息
func (*WsWorkerInfo) WsWorkerNewInfo(D model.WsProtocol, Ws websocket.Connection) (err error) {
	//检查现有workerList有无信息
	if model.WorkerWsList == nil {
		global.Log.Info("初始化WorkerWsList")
		model.WorkerWsList = make(map[string]model.Worker)
	}
	worker := model.WorkerWsList[Ws.ID()]
	if worker != (model.Worker{}) {
		return errors.New("已存在该连接信息")
	}

	//没有的话新建
	err = json.Unmarshal([]byte(D.Body), &worker)
	if err != nil {
		return err
	}

	global.Log.Info("ID"+Ws.ID(), worker)
	model.WorkerWsList[Ws.ID()] = worker
	return nil
}

// WsWorkerUpdateState 更新客户端状态
func (*WsWorkerInfo) WsWorkerUpdateState(D model.WsProtocol, Ws websocket.Connection) (err error) {
	//worker准备好指的是IP准备好,打开了代理IP端口
	//修改客户端Map状态
	worker := model.WorkerWsList[Ws.ID()]
	if worker == (model.Worker{}) {
		return errors.New("试图更新客户端状态,但是未找到其Worker对应的Map")
	}
	//解析进去
	err = json.Unmarshal([]byte(D.Body), &worker)
	if err != nil {
		return err
	}

	//TODO:校验IP信息
	if worker.Info.InternetIp == "" {

	}
	//TODO:加入IP池

	//更新状态
	global.Log.Info("更新状态,ID"+Ws.ID(), worker)
	model.WorkerWsList[Ws.ID()] = worker

	//触发下发任务函数
	return nil
}

// WsWorkerGetTask 客户端申请一个任务
func VerifyInternetIP(D model.WsProtocol, Ws websocket.Connection) (err error) {
	//ToDo:Redis里面List,Task任务号与真实的东西串联起来,检查链表,如果有,则IP重复,否则可用
	//检查数据库任务情况,
	return nil
}
