package model

import (
	"io"
	"log"
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
	addr := strings.Split(r.RemoteAddr, ":")
	log.Println(addr[0], r.Method, r.RequestURI, http.StatusOK)

}
