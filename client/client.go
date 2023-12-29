package client

import (
	"fmt"
	"io"
	"os"
	"poster/common/poster_io"
	"poster/common/security"
	"sync"
)

type Client struct {
	file_path   string
	server_name string
	key_name    string
	chunk_size  int
}

var once sync.Once
var client_instance *Client

func GetClientInstance(file_path string, server_name string, key_name string) *Client {
	once.Do(func() {
		client_instance = &Client{
			file_path:   file_path,
			server_name: server_name,
			key_name:    key_name,
			chunk_size:  1024,
		}
	})
	return client_instance
}

func (c *Client) RunPoster() {
	// 打开文件
	f, err := os.Open(c.file_path)
	if err != nil {
		panic("open file error")
	}
	defer f.Close()

	// 初始 - 服务端收到该post后创建文件
	poster_io.PostDataWithGetMethod(5, c.server_name, c.key_name, "", 0, "0")

	// 持续读取到buffer
	buf := make([]byte, c.chunk_size)
	cnt := 0
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic("read file error")
		}
		if err == io.EOF {
			break
		}
		data := buf[:n]
		fmt.Println("recv data", string(data))
		// crypto
		secure_data := security.Encrypt(data)
		// secure_string := string(secure_data)
		fmt.Println("secure string is", secure_data)

		// post
		poster_io.PostDataWithGetMethod(5, c.server_name, c.key_name, secure_data, cnt, "1")

		cnt++
	}
	fmt.Println("file send finished")

	// 终止标记 - 服务端收到后关闭文件
	poster_io.PostDataWithGetMethod(5, c.server_name, c.key_name, "", cnt, "2")

	fmt.Println("run poster")
}
