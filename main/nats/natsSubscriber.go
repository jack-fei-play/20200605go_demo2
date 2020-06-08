package main

import (
	"../coonRedis"
	"../natsStruct"
	"../sqlInsert"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"runtime"
	"strconv"
)

func main() {
	// create server connection
	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("connection fail,", err)
		return
	}
	// subscribe to subject
	natsConnection.Subscribe("ff/nats/demo", func(msg *nats.Msg) {
		//handle the message
		fmt.Printf("received message '%s\n", string(msg.Data)+"'")
		if msg.Data == nil {
			fmt.Printf("message nil")
			return
		}
		FindRedisData(msg.Data)

	})
	//nats
	fmt.Printf("natsSubscriber server success! \r\n")
	// keep the connection alive
	runtime.Goexit()
}

//查询redis获取数据

func FindRedisData(msg []byte) {
	var sendMsg natsStruct.NatsMsg
	json.Unmarshal(msg, &sendMsg)
	for _, natsMsgStruct := range sendMsg.Datas {
		//通过natsMsgStruct.TagId获取redis中数据
		redisDataId := coonRedis.ReadRedis(natsMsgStruct.TagId)
		var natsStruct natsStruct.NatsStruct
		natsStruct.ThisTime = sendMsg.ThisTime
		natsStruct.DeviceId = sendMsg.DeviceId
		natsStruct.Tag = sendMsg.Tag
		atoi, _ := strconv.Atoi(natsMsgStruct.TagId)
		natsStruct.TagId = atoi
		natsStruct.DataId = redisDataId
		natsStruct.Value = natsMsgStruct.Value

		//向mysql插入数据
		sqlInsert.InsertSql(natsStruct)
	}
}
