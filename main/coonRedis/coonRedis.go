package coonRedis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	ip       = "192.168.200.129"
	port     = 6379
	password = 123456
)

type RedisStruct struct {
	data_id string
}

//往redis写入数据
func WriteRedis() {

	conn, err := redis.Dial("tcp", "192.168.1.124:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("redis connection fail!", err)
		return
	}
	fmt.Println("redis connection success")
	//map中数据
	redisMap := make(map[string]RedisStruct)
	redisMap["ff-nats:tag_id:1101"] = RedisStruct{"1201"}
	redisMap["ff-nats:tag_id:1102"] = RedisStruct{"1202"}
	redisMap["ff-nats:tag_id:1103"] = RedisStruct{"1203"}
	for key := range redisMap {
		//循环写入
		_, err = conn.Do("hmset", key, "data_id", redisMap[key].data_id)
		if err != nil {
			fmt.Println("redis insert fail!", err)
			return
		}
	}
	fmt.Println("redis write success!")
	defer conn.Close()
}

//从redis读出数据
func ReadRedis(id string) string {
	conn, err := redis.Dial("tcp", "192.168.1.124:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("redis connection fail!", err)
		return ""
	}
	fmt.Println("redis connection success")
	//读取
	s := "ff-nats:tag_id:" + id
	fmt.Println(s)
	do, err := conn.Do("hget", "ff-nats:tag_id:"+id, "data_id")
	//do, err := redis.Values(conn.Do("HGETALL", "ff-nats:tag_id:" + id))
	if err != nil {
		fmt.Println("redis insert fail!", err)
		return ""
	}
	//var redisStruct RedisStruct
	s2 := do.([]byte)
	redisDataId := string(s2[:])
	return redisDataId
	//var redisStruct RedisStruct
	//redisStruct.data_id =s3
	//fmt.Println("redis read success!")
	//defer conn.Close()
	//return redisStruct
}

//s2 := do.([]byte)
//s3 := string(s2[:])
//fmt.Println(s3)
