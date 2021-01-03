package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	var resp string

	// 建立连接
	client, err := rpc.Dial("tcp", ":8087")
	if err != nil {
		log.Fatalln("rpc.DialHTTP error:", err)
	}
	reader := bufio.NewReader(os.Stdin)
	//
	for {
		log.Println("-> ")
		// 读取输入
		req, _ := reader.ReadString('\n')
		req = strings.Trim(req, "\r\n")
		if req == "exit" {
			return
		}
		// 同步的方式调用rpc()
		if client.Call("EncryptionUtil.Md5", req, &resp) != nil {
			log.Fatalf("client.Call error: %v", err)
		}
		log.Println("加密值:", resp)
	}
}
