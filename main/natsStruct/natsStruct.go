package natsStruct

type NatsStruct struct { //数据库存储信息的结构体
	Id       int     //'自己生成随机id'
	ThisTime uint64  //'当前时间'
	DeviceId uint64  //'设备Id号'
	Tag      string  //'标签信息'
	DataId   string  //'dataId'
	TagId    int     //'tagId'
	Value    float32 //'该条数据上传的值'
}
type NatsMsg struct { //nats发送消息结构体
	ThisTime uint64        `json:"this_time"`
	DeviceId uint64        `json:"device_id"`
	Tag      string        `json:"tag"`
	Datas    []NatsMsgData `json:"datas"`
}
type NatsMsgData struct { //nats发送消息结构体
	TagId string  `json:"tag_id"`
	Value float32 `json:"data_id"`
}
