package natsStruct

type NatsStruct struct {
	Id       int    //'自己生成随机id'
	ThisTime int    //'当前时间'
	DeviceId string //'设备Id号'
	Tag      string //'标签信息'
	DataId   int    //'dataId'
	TagId    int    //'tagId'
	Value    int    //'该条数据上传的值'
}
