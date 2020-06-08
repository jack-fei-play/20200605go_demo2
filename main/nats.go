package main

import "../main/coonRedis"

func main() {
	//向redis写入数据
	coonRedis.WriteRedis()
	//监听nats消息，获取内容后从redis中读取
	//slice1 := make([] string)
	//var slice = []string{"ff-nats:tag_id:1101", "ff-nats:tag_id:1102", "ff-nats:tag_id:1103"}
	////strings := append(slice, "ff-nats:tag_id:1101", "ff-nats:tag_id:1102", "ff-nats:tag_id:1103")
	//redis.ReadRedis(slice)
	//sqlInsert.InsertSql()
}
