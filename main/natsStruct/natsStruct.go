package natsStruct

type NatsStruct struct {
	Id       int     //'自己生成随机id'
	ThisTime uint64  //'当前时间'
	DeviceId uint64  //'设备Id号'
	Tag      string  //'标签信息'
	DataId   string  //'dataId'
	TagId    int     //'tagId'
	Value    float32 //'该条数据上传的值'
}
type NatsMsg struct {
	ThisTime uint64        `json:"this_time"`
	DeviceId uint64        `json:"device_id"`
	Tag      string        `json:"tag"`
	Datas    []NatsMsgData `json:"datas"`
}
type NatsMsgData struct {
	TagId string  `json:"tag_id"`
	Value float32 `json:"data_id"`
}
