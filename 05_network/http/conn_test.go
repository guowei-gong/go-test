package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	// 监听一个未被占用的端口
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	// 启动 http 服务
	go http.Serve(ln, nil)

	// 发起一个 Get 请求
	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "Hello World" {
		t.Fatal("expected Hello World, but got", string(body))
	}
}
