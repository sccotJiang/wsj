package main

import (
	"github.com/sccotJiang/wsj/internal/entities/namespaces"
	"github.com/sccotJiang/wsj/internal/websocket/server"
	"github.com/sccotJiang/wsj/internal/websocket/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func init() {
	go listenSyscall()
	//初始化消息监听逻辑
	service.InitClientManager()
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("asdasd")
		server.Connect(r, w)
	})

	cm := service.GetClientManger()
	http.HandleFunc("/ws/caller", func(w http.ResponseWriter, r *http.Request) {
		log.Println("asdasd")
		server.CreateConnect(r, w, cm)
	})

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Println("listen fail")
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

//监听系统信号
func listenSyscall() {
	sigc := make(chan os.Signal, 1)
	//监听和捕获信号量。首先定义一个c，用于传递信号量，然后指定哪些信号量是需要被捕获的，如果不指定，就会捕获任何信号量。
	//没有捕获到信号量时就一直阻塞
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	select {
	case sig := <-sigc:
		log.Printf("receive system signal %v, start terminating", sig)
		wg := sync.WaitGroup{}
		for _, n := range namespaces.Namespaces {
			wg.Add(1)
			go service.GraceTerminate(n, &wg)
		}
		wg.Wait()
		log.Println("finish terminating")
		os.Exit(3)
	}
}
