package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	ip   = "192.168.200.129"
	port = 6379
)

type RedisStruct struct {
	data_id int
}

//往redis写入数据
func WriteRedis() {

	conn, err := redis.Dial("tcp", "192.168.200.129:6379")
	if err != nil {
		fmt.Println("redis connection fail!", err)
		return
	}
	fmt.Println("redis connection success")
	//map中数据
	redisMap := make(map[string]RedisStruct)
	redisMap["ff-nats:tag_id:1101"] = RedisStruct{1201}
	redisMap["ff-nats:tag_id:1102"] = RedisStruct{1202}
	redisMap["ff-nats:tag_id:1103"] = RedisStruct{1203}
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

//往redis写入数据
func ReadRedis(arr []string) {

	conn, err := redis.Dial("tcp", "192.168.200.129:6379")
	if err != nil {
		fmt.Println("redis connection fail!", err)
		return
	}
	fmt.Println("redis connection success")

	for _, value := range arr {
		//循环写入
		do, err := redis.Values(conn.Do("HGETALL", value))
		if err != nil {
			fmt.Println("redis insert fail!", err)
		} else {
			for _, v := range do {
				fmt.Printf("%s ", v.([]byte))
			}
		}

	}
	fmt.Println("redis read success!")
	defer conn.Close()
}
