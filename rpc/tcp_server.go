package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// 加密工具类
type EncryptionUtil struct {
}

// 加密方法
func (eu *EncryptionUtil) Md5(req string, resp *string) error {
	*resp = fmt.Sprintf("%x", md5.Sum([]byte(req)))
	return nil
}
func main() {
	// 功能对象注册
	if err := rpc.Register(new(EncryptionUtil)); err != nil {
		log.Fatalln("rpc.Register error: ", err)
	}
	// 端口监听
	listen, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalln("net.Listen error: ", err)
	}
	// 启动服务
	log.Println("服务已启动")
	rpc.Accept(listen)
}
