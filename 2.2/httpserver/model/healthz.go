package model

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Healthz(w http.ResponseWriter, r *http.Request) {

	//request header写入response header
	header := r.Header
	for k, _ := range header {
		w.Header().Set(k, header.Get(k))
	}

	//读取系统变量获取VERSION
	var VERSION string
	VERSION = os.Getenv("VERSION")
	w.Header().Set("VERSION", VERSION)

	io.WriteString(w, "200")

	//打印日志
	addr := strings.Split(r.RemoteAddr, ":")
	fmt.Println(QueryTime(), addr[0], r.Method, r.RequestURI, http.StatusOK)

}
