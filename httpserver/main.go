package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"httpserver/metrics"
	"httpserver/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	fmt.Println("Starting http server...")
	metrics.Register()

	mux := http.NewServeMux()
	mux.HandleFunc("/", model.Index)
	mux.HandleFunc("/healthz", model.Healthz)
	mux.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	// make sure idle connections returned
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		//监听 Ctrl+C 信号
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); nil != err {
			log.Fatalf("server shutdown failed, err: %v\n", err)
		}

		time.Sleep(time.Second * 10)
		log.Println("server gracefully shutdown")

		close(processed)
	}()

	// serve
	err := s.ListenAndServe()
	if http.ErrServerClosed != err {
		log.Fatalf("server not gracefully shutdown, err :%v\n", err)
	}

	// waiting for goroutine above processed
	<-processed

}
