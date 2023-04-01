package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/current_time", sayHello)
	mux.HandleFunc("/v2/current_time", sayHelloV2)

	// 创建服务器
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	// 监听端口并提供服务
	log.Println("starting httpserver at :8000")
	log.Fatal(server.ListenAndServe())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, current time: " + time.Now().String()))
}

func sayHelloV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO4, current time: " + time.Now().Format("2006-01-02 15:04:05")))
}
