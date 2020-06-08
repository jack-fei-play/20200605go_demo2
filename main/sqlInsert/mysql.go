package sqlInsert

import (
	"../natsStruct"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"log"
	"math/rand"
	"strings"
	"time"
)

func InsertSql(natsStruct natsStruct.NatsStruct) {
	//读取ini文件
	host, port, username, password, database := readIni()
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8"}, "")
	var DB *sql.DB
	DB, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T", DB)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
	//没使用预编译sql对象，可能引发sql注入
	//number := randomNumber()
	//id := strconv.Itoa(number)
	//dataId:= natsStruct.DataId
	//tagId:= strconv.Itoa(natsStruct.TagId)
	//value:= strconv.FormatFloat(float64(natsStruct.Value), 'E', -1, 32)
	//thisTime:= strconv.Itoa(int(natsStruct.ThisTime))
	//deviceId:= strconv.Itoa(int(natsStruct.DeviceId))
	//sql := "insert into `ff-nats-demo` values ("+id+","+thisTime+",\""+deviceId+"\","+"\""+natsStruct.Tag+"\","+dataId+","+tagId+","+value+")"
	//result, err := DB.Exec(sql)
	//if err != nil {
	//	fmt.Println("sql执行出错！",err)
	//}
	//DB.Close()
	//fmt.Println("%T", result)
	//fmt.Println("insert end")
	sql := "insert into `ff-nats-demo` values (?,?,?,?,?,?,?)"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	result, err := stmt.Exec(randomNumber(), natsStruct.ThisTime, natsStruct.DeviceId, natsStruct.Tag, natsStruct.DataId, natsStruct.TagId, natsStruct.Value)
	if err != nil {
		fmt.Println(err)
	}
	DB.Close()
	fmt.Println("%T", result)
	fmt.Println("insert end")

}
func readIni() (host string, port string, username string, password string, database string) {
	//读取ini文件
	cfg, err := ini.Load("config.ini")
	//if err != nil {
	//	fmt.Println("load config.ini fail,",err)
	//}
	getErr("load config", err)
	// 获取mysql分区的key
	fmt.Println(cfg.Section("mysql").Key("host").String()) // 将结果转为string
	fmt.Println(cfg.Section("mysql").Key("port").Int())    // 将结果转为int
	host = cfg.Section("mysql").Key("host").String()
	port = cfg.Section("mysql").Key("port").String()
	username = cfg.Section("mysql").Key("username").String()
	password = cfg.Section("mysql").Key("password").String()
	database = cfg.Section("mysql").Key("database").String()
	return host, port, username, password, database
}

//config.ini获取错误处理
func getErr(msg string, err error) {
	if err != nil {
		log.Printf("%v err->%v\n", msg, err)
	}
}

//randomNumber
func randomNumber() int {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)
	return rnd
}
