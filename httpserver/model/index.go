package model

import (
	"httpserver/metrics"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)

}

func Index(w http.ResponseWriter, r *http.Request) {

	timer := metrics.NewTimer()
	defer timer.ObserverTotal()

	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))

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
