package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func TestConn(t *testing.T) {
	// 模拟请求对象
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	// 模拟响应对象
	w := httptest.NewRecorder()
	helloHandler(w, req)

	bytes, _ := ioutil.ReadAll(w.Result().Body)
	if string(bytes) != "Hello World" {
		t.Fatal("expected Hello World, but got", string(bytes))
	}
}
