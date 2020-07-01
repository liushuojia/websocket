package main

import (
	"log"
	"websocket/router"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	log.Println("[gin]", "start server")

	//进程停止时候运行
	ch := make(chan os.Signal, 1)
	signal.Notify(
		ch,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	go func() {
		s := <-ch
		log.Println("[gin] 停止服务", s)
		os.Exit(1)
	}()

	// http server
	router.HttpStart()
	return
}
