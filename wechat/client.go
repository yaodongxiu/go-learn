package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9601")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan string)
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- "ok"
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
