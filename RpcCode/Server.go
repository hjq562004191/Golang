package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

func (ma *MathUtil) CalculateCArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}

func main() {
	//初始化一个指针类型
	mathUtil := new(MathUtil)

	//对服务对象进行注册
	err := rpc.Register(mathUtil)

	if err != nil {
		panic(err.Error())
	}

	//使远程调用者以HTTP的形式调用方法
	rpc.HandleHTTP()

	//设置特定端口进行监听
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err.Error())
	}

	http.Serve(listener, nil)
}
