package main

import (
	"flag"
	"fmt"
	"github.com/sccotJiang/wsj/internal/entities/namespaces"
	"github.com/sccotJiang/wsj/internal/websocket/server"
	"github.com/sccotJiang/wsj/internal/websocket/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func init() {
	go listenSyscall()
	//定义命令行参数对应的变量
	var cliName = flag.String("name", "默认姓名", "输入你的姓名")
	flag.Parse() //把用户传递的命令行参数解析为对应变量的值 go run main.go -name=1212
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *cliName)
	//初始化环境
	server.InitEnv()
	//初始化组件
	server.InitInternalServer()
}

func main() {
	http.HandleFunc("/ws/caller", func(w http.ResponseWriter, r *http.Request) {
		server.InternalManager{}
	})
	<-time.After(time.Second * 10)
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
