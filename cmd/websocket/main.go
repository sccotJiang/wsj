package websocket

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go listenSyscall()
}
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
	}
}
