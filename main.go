package main

import (
	"flag"
	"fmt"
	"poster/client"
	"poster/server"
)

func main() {

	var t int
	var p string
	var s string
	var k string
	flag.IntVar(&t, "t", 0, "运行环境 0: 服务端, 1: 客户端")
	flag.StringVar(&p, "p", "", "文件路径")
	flag.StringVar(&s, "s", "", "服务器路径")
	flag.StringVar(&k, "k", "", "文件key")
	flag.Parse()
	fmt.Println(t, p, s)
	if t == 1 {
		println("run client")
		c := client.GetClientInstance(p, s, k)
		c.RunPoster()
	} else {
		println("run server")
		s := server.GetServerInstance("0.0.0.0", 9999)
		s.Run()
	}
}
