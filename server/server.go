package server

import (
	"fmt"
	"net/http"
	"os"
	"poster/common/security"
	"sync"
)

type Server struct {
	port         int
	listen       string
	file_manager map[string](*os.File)
}

var once sync.Once
var server_instance *Server

func GetServerInstance(listen_ip string, port int) *Server {
	once.Do(func() {
		server_instance = &Server{
			port:         port,
			listen:       listen_ip,
			file_manager: make(map[string](*os.File)),
		}
	})
	return server_instance
}

func (s *Server) Run() {
	fmt.Println("server start and listen ip:port")
	http.HandleFunc("/", s.dataHandler)
	listen_str := fmt.Sprintf("%s:%d", s.listen, s.port)
	err := http.ListenAndServe(listen_str, nil)
	if err != nil {
		panic("server listen error")
	}
}

func (s *Server) dataHandler(c http.ResponseWriter, req *http.Request) {
	fmt.Println("\nhandle")
	req.ParseForm()
	key := req.FormValue("key")
	data := req.FormValue("data")
	cmd := req.FormValue("cmd")
	cnt := req.FormValue("cnt")
	fmt.Println("server get: ", key, data, cnt)
	// fmt.Println("key=", key, "\ndata=\n", data)
	write_path := "target/" + key
	if cmd == "0" {
		// 创建文件 - 阻塞
		f, err := os.Create(write_path)
		if err != nil {
			panic("create file error")
		}
		s.file_manager[key] = f
	} else if cmd == "2" {
		// 关闭文件
		s.file_manager[key].Close()
		fmt.Println("file closed")
		delete(s.file_manager, key)
	} else {
		// 持续写入 - 阻塞写
		decode_data := security.Decrypt(data)
		_, err := s.file_manager[key].Write(decode_data)
		if err != nil {
			panic("write file failed")
		}
	}

}
