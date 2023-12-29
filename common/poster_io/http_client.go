package poster_io

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// type HttpClient struct {
// 	hc      http.Client
// }

var http_client_instance *http.Client
var once sync.Once

func GetHttpClientInstance(timeout time.Duration) *http.Client {
	once.Do(func() {
		http_client_instance = &http.Client{Timeout: timeout * time.Second}
	})
	return http_client_instance
}

func PostDataWithGetMethod(timeout time.Duration, url string, key string, data string, cnt int, cmd string) {
	c := GetHttpClientInstance(timeout)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("http new request err")
	}
	q := req.URL.Query()
	q.Add("key", key)
	q.Add("data", data)
	q.Add("cnt", fmt.Sprintf("%d", cnt))
	q.Add("cmd", cmd)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("http req failed")
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("http success")
	}

}
