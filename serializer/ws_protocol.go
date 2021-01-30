package serializer

import (
	"encoding/json"
	"fmt"
	"leopard/model"
	"time"
)

//WsSerializerReq 序列化
func WsSerializerReq(d model.WsProtocol) []byte {
	d.Time = time.Now().Unix()
	jsonByte, _ := json.Marshal(d)
	return jsonByte
}

//WsDeserialization 反序列化
func WsDeserialization(data []byte) (m *model.WsProtocol, err error) {
	fmt.Println(string(data))
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
