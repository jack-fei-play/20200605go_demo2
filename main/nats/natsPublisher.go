package main

import (
	"../natsStruct"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	// create server connection and defer close
	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("connection fail,", err)
		return
	}
	defer natsConnection.Close()
	fmt.Println("connected to " + nats.DefaultURL)
	//sendMsg对象信息封装
	var sendMsg natsStruct.NatsMsg
	sendMsg.DeviceId = 897979799
	sendMsg.Tag = "dtu0000001"
	sendMsg.ThisTime = 1500
	sendMsg.Datas = []natsStruct.NatsMsgData{
		{
			TagId: "1101",
			Value: 1.3,
		},
		{
			TagId: "1102",
			Value: 25,
		},
		{
			TagId: "1103",
			Value: 80,
		},
	}
	sendJson, err := json.Marshal(sendMsg)
	subject := "ff/nats/demo"
	//nats发送消息
	natsConnection.Publish(subject, sendJson)
	fmt.Printf("published message on subject " + subject)
	s := string(sendJson)
	fmt.Println("send message info:\r\n", s)
	time.Sleep(30 * time.Second)
}
