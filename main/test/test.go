package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

//config.ini获取错误处理
func getErr(msg string, err error) {
	if err != nil {
		log.Printf("%v err->%v\n", msg, err)
	}
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

func main() {
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

	//sql := "insert into `ff-nats-demo` values (3213,1,\"897979799\",\"dtu0000001\",1.3,1101,1.3)"
	sql := "insert into `ff-nats-demo` values (?,?,?,?,?,?,?)"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		fmt.Println("sql执行出错！", err)
	}
	result, err := stmt.Exec(911, 1, "897979799", "dtu0000001", 1.3, 1101, 1.3)
	DB.Close()
	fmt.Println("%T", result)
	fmt.Println("insert end")

}
