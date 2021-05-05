// netcat 是一个简单的TCP服务器读/写客户端
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9601")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan int, 1)
	go func() {
		io.Copy(os.Stdout, conn) // 注意：忽略错误
		log.Println("done")
		done <- 1 // 向主Goroutine发出信号
	}()
	mustCopy(conn, os.Stdin)
	<-done // 等待后台goroutine完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
