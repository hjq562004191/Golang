package main

import (
	"Golang/param"
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}

	//圆的半径
	//var req float32
	//req = 3

	//两数相加
	var a1 float32
	var a2 float32
	a1 = 2
	a2 = 3
	param := param.Addnum{a1, a2}
	var resp *float32
	err = client.Call("MathUtil.Add", param, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)
}
