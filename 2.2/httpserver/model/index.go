package model

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {

	//解决访问出现两次的问题
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}

	//request header写入response header
	header := r.Header
	for k, _ := range header {
		w.Header().Set(k, header.Get(k))
	}

	//读取系统变量获取VERSION
	var VERSION string
	VERSION = os.Getenv("VERSION")
	w.Header().Set("VERSION", VERSION)

	io.WriteString(w, "Hi Teacher, This is my job!")

	//打印日志
	addr := getCurrentIp(r)
	fmt.Println(QueryTime(), addr[0], r.Method, r.RequestURI, http.StatusOK)

}

func getCurrentIp(r *http.Request) string {

	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	return ip

}
