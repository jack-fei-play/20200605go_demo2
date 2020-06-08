package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//max := new(big.Int).Lsh(big.NewInt(1), 128)
	//fmt.Println(max)
	//// 生成大整数随机数
	//serialNumber, _ := rand.Int(rand.Reader, max)
	//fmt.Println(serialNumber)
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)
	fmt.Println(rnd)

}
