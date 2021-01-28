package serializer

import (
	"encoding/json"
	"leopard/model"
	"time"
)

//WsSerializerReq 序列化
func WsSerializerReq(commandType string, data string) []byte {
	jsonByte, _ := json.Marshal(model.WsProtocol{
		CommandType: commandType,
		Body:        data,
		Time:        time.Now().Unix(),
	})
	return jsonByte
}

//WsDeserialization 反序列化
func WsDeserialization(data []byte) (m model.WsProtocol, err error) {
	err = json.Unmarshal(data, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}

//WSTransition 纯粹的中转,它做的事情就是将进来的包序列化一下,然后扔出去
func WSTransition(res model.WsProtocol) []byte {
	jsonByte, _ := json.Marshal(res)
	return jsonByte
}
